### TCMalloc简介

由于Go的内存分配采用的是TCMalloc算法，所以在这之前，咱们得先了解一下TCMalloc的具体原理，可以参考一下这篇文章[TCMalloc : Thread-Caching Malloc](http://goog-perftools.sourceforge.net/doc/tcmalloc.html)。总的来说，这个算法有以下三个优点：

1. 分配更快
2. 在多线程的情况下减少了锁的竞争
3. 小对象的高利用率

TCMalloc给每一个线程都有分配一个线程本地缓存thread-local cache，它能够满足小内存块的分配。内存块也能够在其需要的时候从central移动到thread-local cache中，并且也会周期性的从thread-local cache合并到central中，方便回收后给其他线程进行使用。在分配空间大于32K时，会直接从heap中进行分配空间。

咱们在这里可以将其看成多级缓存来进行理解。如下图所示：

![](http://images.hcyhj.cn/blogimages/mallocgc/tcmalloc.png)

如上图，在分配一个新的small内存块(<32K)时，先从thread-local中去获取，因为是从当前线程缓存中去分配空间，所以不会出现并发问题，因而不需要上锁，效率相对来说比较高。如果在thread-local没有足够的空间，则需要向上级central进行申请，这里由于是多个thread共享的操作，所以要加上锁来避免并发问题。central没有足够的空间，则需要向heap进行申请，同样，heap是一个application共享的，也需要加上锁，然后会把申请到的空间逐级返回，并进行缓存。



了解了一下TCMalloc的相关知识点后，咱们进行Go的内存分配相关源码的研究。

### mallocgc

```go
func mallocgc(size uintptr, typ *_type, needzero bool) unsafe.Pointer {
	if gcphase == _GCmarktermination {
		throw("mallocgc called with gcphase == _GCmarktermination")
	}
	// 对0尺寸的对象进行快速分配，返回的都是同一个全局对象
	if size == 0 {
		return unsafe.Pointer(&zerobase)
	}

	// _GCoff=0，GC没有运行
    // _GCmark=1，GC正在进行标记工作
    // _GCmarktermination=2，GC已经结束
	var assistG *g
	if gcBlackenEnabled != 0 {
		// Charge the current user G for this allocation.
		assistG = getg()
		if assistG.m.curg != nil {
			assistG = assistG.m.curg
		}
		// Charge the allocation against the G. We'll account
		// for internal fragmentation at the end of mallocgc.
		assistG.gcAssistBytes -= int64(size)

		if assistG.gcAssistBytes < 0 {
			// This G is in debt. Assist the GC to correct
			// this before allocating. This must happen
			// before disabling preemption.
			gcAssistAlloc(assistG)
		}
	}

	// Set mp.mallocing to keep from being preempted by GC.
    // 标记该线程正在分配内存，防止被GC进行抢占
	mp := acquirem()
	mp.mallocing = 1

	shouldhelpgc := false
	dataSize := size
    // 获取线程本地缓存
	c := gomcache()
	var x unsafe.Pointer
	noscan := typ == nil || typ.kind&kindNoPointers != 0
	if size <= maxSmallSize {
		if noscan && size < maxTinySize {
			// 小分配器：
            // 可以将多个小的分配请求分配到一个内存块上，这块内存在所有分配在
            // 该block上的object均不可达之后才会被释放掉，这些object必须是
            // 非指针类型的，这样可以保证潜在的内存浪费会最小化。
			off := c.tinyoffset
			// 进行字节对齐
			if size&7 == 0 {
				off = round(off, 8)
			} else if size&3 == 0 {
				off = round(off, 4)
			} else if size&1 == 0 {
				off = round(off, 2)
			}
			if off+size <= maxTinySize && c.tiny != 0 {
				// 上一次分配的tiny block还没被用完，如果此次申请的内存大小合适
                // 则直接在tiny block上继续分配即可
				x = unsafe.Pointer(c.tiny + off)
				c.tinyoffset = off + size
				c.local_tinyallocs++
				mp.mallocing = 0
				releasem(mp)
				return x
			}
			// tinySpanClass值为5，分配一个新的16字节的span
			span := c.alloc[tinySpanClass]
            // 进行快速分配，从当前的mcache中进行分配
			v := nextFreeFast(span)
			if v == 0 {
                // 这里有可能从central甚至是heap中去分配内存块
				v, _, shouldhelpgc = c.nextFree(tinySpanClass)
			}
            // 将分配的16字节置为空
			x = unsafe.Pointer(v)
			(*[2]uint64)(x)[0] = 0
			(*[2]uint64)(x)[1] = 0
			if size < c.tinyoffset || c.tiny == 0 {
                // 当此次分配之后还有剩余空间时，替换old tiny block
				c.tiny = uintptr(x)
				c.tinyoffset = size
			}
			size = maxTinySize
		} else {
			var sizeclass uint8
            // 计算出一个合适的span size
			if size <= smallSizeMax-8 {
				sizeclass = size_to_class8[(size+smallSizeDiv-1)/smallSizeDiv]
			} else {
				sizeclass = size_to_class128[(size-smallSizeMax+largeSizeDiv-1)/largeSizeDiv]
			}
			size = uintptr(class_to_size[sizeclass])
            // 得到具体span的索引
			spc := makeSpanClass(sizeclass, noscan)
			span := c.alloc[spc]
            // 进行快速分配
			v := nextFreeFast(span)
			if v == 0 {
                // 当前尺寸的span已经用完了，可能需要重新填充
                // 这里有可能从central甚至是heap中去分配内存块
				v, span, shouldhelpgc = c.nextFree(spc)
			}
			x = unsafe.Pointer(v)
			if needzero && span.needzero != 0 {
				memclrNoHeapPointers(unsafe.Pointer(v), size)
			}
		}
	} else {
		var s *mspan
		shouldhelpgc = true
		systemstack(func() {
            // 进行大size的分配,直接从heap中进行分配
			s = largeAlloc(size, needzero, noscan)
		})
		s.freeindex = 1
		s.allocCount = 1
		x = unsafe.Pointer(s.base())
		size = s.elemsize
	}

	var scanSize uintptr
	if !noscan {
		// If allocating a defer+arg block, now that we've picked a malloc size
		// large enough to hold everything, cut the "asked for" size down to
		// just the defer header, so that the GC bitmap will record the arg block
		// as containing nothing at all (as if it were unused space at the end of
		// a malloc block caused by size rounding).
		// The defer arg areas are scanned as part of scanstack.
		if typ == deferType {
			dataSize = unsafe.Sizeof(_defer{})
		}
        // 
		heapBitsSetType(uintptr(x), size, dataSize, typ)
		if dataSize > typ.size {
			// Array allocation. If there are any
			// pointers, GC has to scan to the last
			// element.
			if typ.ptrdata != 0 {
				scanSize = dataSize - typ.size + typ.ptrdata
			}
		} else {
			scanSize = typ.ptrdata
		}
		c.local_scan += scanSize
	}

	// Ensure that the stores above that initialize x to
	// type-safe memory and set the heap bits occur before
	// the caller can make x observable to the garbage
	// collector. Otherwise, on weakly ordered machines,
	// the garbage collector could follow a pointer to x,
	// but see uninitialized memory or stale heap bits.
	publicationBarrier()

	// Allocate black during GC.
	// All slots hold nil so no scanning is needed.
	// This may be racing with GC so do it atomically if there can be
	// a race marking the bit.
	if gcphase != _GCoff {
		gcmarknewobject(uintptr(x), size, scanSize)
	}
	mp.mallocing = 0
	releasem(mp)

	if rate := MemProfileRate; rate > 0 {
		if rate != 1 && int32(size) < c.next_sample {
			c.next_sample -= int32(size)
		} else {
			mp := acquirem()
			profilealloc(mp, x, size)
			releasem(mp)
		}
	}

	if assistG != nil {
		// Account for internal fragmentation in the assist
		// debt now that we know it.
		assistG.gcAssistBytes -= int64(size - dataSize)
	}

	if shouldhelpgc {
		if t := (gcTrigger{kind: gcTriggerHeap}); t.test() {
			gcStart(t)
		}
	}

	return x
}
```



### makeSpanClass

下图是`mcache`中alloc数组对应的结构，一共有67种规格的span,然后有scan和noscan两种类型，则一共是134个。

![](http://images.hcyhj.cn/blogimages/mallocgc/mcache_alloc.png)

```go
// 如果noscan = true
// 例如sizeClass=2,那么对于的spanClass = 2*2 + 1 = 5
// 上述例子对应16字节的noscan类型的span
func makeSpanClass(sizeclass uint8, noscan bool) spanClass {
	return spanClass(sizeclass<<1) | spanClass(bool2int(noscan))
}
```



### mspan.nextFreeFast

```go
func nextFreeFast(s *mspan) gclinkptr {
    // 获取allocCache最小的一个值为1的bit位的索引
    // 例如 ..11101000,此时计算出的索引位为3
	theBit := sys.Ctz64(s.allocCache) // Is there a free object in the allocCache?
	if theBit < 64 {
		result := s.freeindex + uintptr(theBit)
		if result < s.nelems {
			freeidx := result + 1
			if freeidx%64 == 0 && freeidx != s.nelems {
                // 由于下一个freeidx已经不在当前allocCache中了，这里要对allocCache
                // 进行refill,具体逻辑在nextFreeIndex方法中
				return 0
			}
			s.allocCache >>= uint(theBit + 1)
			s.freeindex = freeidx
			s.allocCount++
			return gclinkptr(result*s.elemsize + s.base())
		}
	}
	return 0
}
```



### mspan.nextFreeIndex

```go
func (s *mspan) nextFreeIndex() uintptr {
	sfreeindex := s.freeindex
	snelems := s.nelems
	if sfreeindex == snelems {
		return sfreeindex
	}
	if sfreeindex > snelems {
		throw("s.freeindex > s.nelems")
	}

	aCache := s.allocCache
	// 计算aCache中bit位为1的位置的索引
	bitIndex := sys.Ctz64(aCache)
	for bitIndex == 64 {
		// Move index to start of next cached bits.
        // 比如咱们的sfreeindex为0，则下一个sfreeindex为64
		sfreeindex = (sfreeindex + 64) &^ (64 - 1)
		if sfreeindex >= snelems {
			s.freeindex = snelems
			return snelems
		}
		whichByte := sfreeindex / 8
		// 重新用下一个64位去填充
		s.refillAllocCache(whichByte)
		aCache = s.allocCache
		bitIndex = sys.Ctz64(aCache)
		// 没有可用的缓存bit，获取接下来的8个字节的标记位并进行重试
	}
	result := sfreeindex + uintptr(bitIndex)
	if result >= snelems {
		s.freeindex = snelems
		return snelems
	}
	// 可以一次性偏移多个位置，例如
    // 当前allocCache为11011000，经过sys.Ctz64计算后得到最小的bitIndex为3
    // 此时会直接将allocCache向左偏移4位变成00001101，一次性将前面位数为0的全部排除掉了
	s.allocCache >>= uint(bitIndex + 1)
	sfreeindex = result + 1

	if sfreeindex%64 == 0 && sfreeindex != snelems {
		// 进行预填充，此时sfreeindex已经是下一个allocCache的起点了
        // 这里需要将下一个64bit位填充到allocCache中
		whichByte := sfreeindex / 8
		s.refillAllocCache(whichByte)
	}
	s.freeindex = sfreeindex
	return result
}
```



```go
func (s *mspan) refillAllocCache(whichByte uintptr) {
    // 从当前位置开始，往后读取8个字节长度（64bits），并经过相应的运算填充到allocCache中
	bytes := (*[8]uint8)(unsafe.Pointer(s.allocBits.bytep(whichByte)))
	aCache := uint64(0)
	aCache |= uint64(bytes[0])
	aCache |= uint64(bytes[1]) << (1 * 8)
	aCache |= uint64(bytes[2]) << (2 * 8)
	aCache |= uint64(bytes[3]) << (3 * 8)
	aCache |= uint64(bytes[4]) << (4 * 8)
	aCache |= uint64(bytes[5]) << (5 * 8)
	aCache |= uint64(bytes[6]) << (6 * 8)
	aCache |= uint64(bytes[7]) << (7 * 8)
	s.allocCache = ^aCache
}
```

关注一下`refillAllocCache`这个方法，它会从下一个64位开始计算，进行填充`s.allocCache`这个数组，由于当前64位已经用完了，`s.allocCache`中所有的bit都是0（当其中有bit位为1时，Ctz64方法便会返回其索引）。而`s.allocBits`中刚好gc过后还存在的标记位为1，被回收的标记位为0，这里进行运算（将0置位1,1置位0），将`s.allocCache`进行重新填充，将其可以分配空间的位置标记为1。



### mcache.nextFree

```go
func (c *mcache) nextFree(spc spanClass) (v gclinkptr, s *mspan, shouldhelpgc bool) {
	s = c.alloc[spc]
	shouldhelpgc = false
    // 寻找下一个可用的freeIndex
	freeIndex := s.nextFreeIndex()
	if freeIndex == s.nelems {
		if uintptr(s.allocCount) != s.nelems {
            // 当前span已经被用完了，但是分配的总次数和预算的次数不一致，抛出异常
			println("runtime: s.allocCount=", s.allocCount, "s.nelems=", s.nelems)
			throw("s.allocCount != s.nelems && freeIndex == s.nelems")
		}
        // 从central中获取相应spanClass的span，进行重新填充
		c.refill(spc)
		shouldhelpgc = true
		s = c.alloc[spc]

		freeIndex = s.nextFreeIndex()
	}

	if freeIndex >= s.nelems {
		throw("freeIndex is not valid")
	}

	v = gclinkptr(freeIndex*s.elemsize + s.base())
	s.allocCount++
	if uintptr(s.allocCount) > s.nelems {
		println("s.allocCount=", s.allocCount, "s.nelems=", s.nelems)
		throw("s.allocCount > s.nelems")
	}
	return
}
```



### mache.refill

```go
	// sweep generation:
	// if sweepgen == h->sweepgen - 2, the span needs sweeping
	// if sweepgen == h->sweepgen - 1, the span is currently being swept
	// if sweepgen == h->sweepgen, the span is swept and ready to use
	// if sweepgen == h->sweepgen + 1, the span was cached before sweep began and is still cached, and needs sweeping
	// if sweepgen == h->sweepgen + 3, the span was swept and then cached and is still cached
	// h->sweepgen is incremented by 2 after every GC


func (c *mcache) refill(spc spanClass) {
	// Return the current cached span to the central lists.
	s := c.alloc[spc]

	if uintptr(s.allocCount) != s.nelems {
		throw("refill of span with free space remaining")
	}
	if s != &emptymspan {
		// Mark this span as no longer cached.
		if s.sweepgen != mheap_.sweepgen+3 {
			throw("bad sweepgen in refill")
		}
		atomic.Store(&s.sweepgen, mheap_.sweepgen)
	}

	// Get a new cached span from the central lists.
	s = mheap_.central[spc].mcentral.cacheSpan()
	if s == nil {
		throw("out of memory")
	}

	if uintptr(s.allocCount) == s.nelems {
		throw("span has no free space")
	}

	// Indicate that this span is cached and prevent asynchronous
	// sweeping in the next sweep phase.
	s.sweepgen = mheap_.sweepgen + 3

	c.alloc[spc] = s
}
```



### mcentral.cacheSpan

```go
func (c *mcentral) cacheSpan() *mspan {
	// Deduct credit for this span allocation and sweep if necessary.
	spanBytes := uintptr(class_to_allocnpages[c.spanclass.sizeclass()]) * _PageSize
	deductSweepCredit(spanBytes, 0)

	lock(&c.lock)
	sg := mheap_.sweepgen
retry:
	var s *mspan
	for s = c.nonempty.first; s != nil; s = s.next {
		if s.sweepgen == sg-2 && atomic.Cas(&s.sweepgen, sg-2, sg-1) {
			c.nonempty.remove(s)
			c.empty.insertBack(s)
			unlock(&c.lock)
			s.sweep(true)
			goto havespan
		}
		if s.sweepgen == sg-1 {
			// the span is being swept by background sweeper, skip
			continue
		}
		// we have a nonempty span that does not require sweeping, allocate from it
		c.nonempty.remove(s)
		c.empty.insertBack(s)
		unlock(&c.lock)
		goto havespan
	}

	for s = c.empty.first; s != nil; s = s.next {
		if s.sweepgen == sg-2 && atomic.Cas(&s.sweepgen, sg-2, sg-1) {
			// we have an empty span that requires sweeping,
			// sweep it and see if we can free some space in it
			c.empty.remove(s)
			// swept spans are at the end of the list
			c.empty.insertBack(s)
			unlock(&c.lock)
			s.sweep(true)
			freeIndex := s.nextFreeIndex()
			if freeIndex != s.nelems {
				s.freeindex = freeIndex
				goto havespan
			}
			lock(&c.lock)
			// the span is still empty after sweep
			// it is already in the empty list, so just retry
			goto retry
		}
		if s.sweepgen == sg-1 {
			// the span is being swept by background sweeper, skip
			continue
		}
		// already swept empty span,
		// all subsequent ones must also be either swept or in process of sweeping
		break
	}

	unlock(&c.lock)

	// Replenish central list if empty.
	s = c.grow()
	if s == nil {
		return nil
	}
	lock(&c.lock)
	c.empty.insertBack(s)
	unlock(&c.lock)

	// At this point s is a non-empty span, queued at the end of the empty list,
	// c is unlocked.
havespan:
	
    // 计算出freeindex基于allocBits而言的当前基准位置，方便后面填充allocCache
	freeByteBase := s.freeindex &^ (64 - 1)
	whichByte := freeByteBase / 8
	// Init alloc bits cache.
	s.refillAllocCache(whichByte)

	// 调整allocCache以便能够和freeindex对应上
	s.allocCache >>= s.freeindex % 64

	return s
}
```





### largeAlloc

```go
func largeAlloc(size uintptr, needzero bool, noscan bool) *mspan {
	// print("largeAlloc size=", size, "\n")

	if size+_PageSize < size {
		throw("out of memory")
	}
	npages := size >> _PageShift
	if size&_PageMask != 0 {
		npages++
	}

	// Deduct credit for this span allocation and sweep if
	// necessary. mHeap_Alloc will also sweep npages, so this only
	// pays the debt down to npage pages.
	deductSweepCredit(npages*_PageSize, npages)

	s := mheap_.alloc(npages, makeSpanClass(0, noscan), true, needzero)
	if s == nil {
		throw("out of memory")
	}
	s.limit = s.base() + size
	heapBitsForAddr(s.base()).initSpan(s)
	return s
}
```

### mheap.alloc

```go
func (h *mheap) alloc(npage uintptr, spanclass spanClass, large bool, needzero bool) *mspan {
	// Don't do any operations that lock the heap on the G stack.
	// It might trigger stack growth, and the stack growth code needs
	// to be able to allocate heap.
	var s *mspan
	systemstack(func() {
		s = h.alloc_m(npage, spanclass, large)
	})

	if s != nil {
		if needzero && s.needzero != 0 {
			memclrNoHeapPointers(unsafe.Pointer(s.base()), s.npages<<_PageShift)
		}
		s.needzero = 0
	}
	return s
}
```

### mheap.alloc_m

```go
func (h *mheap) alloc_m(npage uintptr, spanclass spanClass, large bool) *mspan {
	_g_ := getg()

	// To prevent excessive heap growth, before allocating n pages
	// we need to sweep and reclaim at least n pages.
	if h.sweepdone == 0 {
		h.reclaim(npage)
	}

	lock(&h.lock)
	// transfer stats from cache to global
	memstats.heap_scan += uint64(_g_.m.mcache.local_scan)
	_g_.m.mcache.local_scan = 0
	memstats.tinyallocs += uint64(_g_.m.mcache.local_tinyallocs)
	_g_.m.mcache.local_tinyallocs = 0

	s := h.allocSpanLocked(npage, &memstats.heap_inuse)
	if s != nil {
		// Record span info, because gc needs to be
		// able to map interior pointer to containing span.
		atomic.Store(&s.sweepgen, h.sweepgen)
		h.sweepSpans[h.sweepgen/2%2].push(s) // Add to swept in-use list.
		s.state = mSpanInUse
		s.allocCount = 0
		s.spanclass = spanclass
		if sizeclass := spanclass.sizeclass(); sizeclass == 0 {
			s.elemsize = s.npages << _PageShift
			s.divShift = 0
			s.divMul = 0
			s.divShift2 = 0
			s.baseMask = 0
		} else {
			s.elemsize = uintptr(class_to_size[sizeclass])
			m := &class_to_divmagic[sizeclass]
			s.divShift = m.shift
			s.divMul = m.mul
			s.divShift2 = m.shift2
			s.baseMask = m.baseMask
		}

		// Mark in-use span in arena page bitmap.
		arena, pageIdx, pageMask := pageIndexOf(s.base())
		arena.pageInUse[pageIdx] |= pageMask

		// update stats, sweep lists
		h.pagesInUse += uint64(npage)
		if large {
			memstats.heap_objects++
			mheap_.largealloc += uint64(s.elemsize)
			mheap_.nlargealloc++
			atomic.Xadd64(&memstats.heap_live, int64(npage<<_PageShift))
		}
	}
	// heap_scan and heap_live were updated.
	if gcBlackenEnabled != 0 {
		gcController.revise()
	}

	//
	unlock(&h.lock)
	return s
}
```

### mheap.allocSpanLocked

```go
func (h *mheap) allocSpanLocked(npage uintptr, stat *uint64) *mspan {
	var s *mspan

	s = h.pickFreeSpan(npage)
	if s != nil {
		goto HaveSpan
	}
	// On failure, grow the heap and try again.
	if !h.grow(npage) {
		return nil
	}
	s = h.pickFreeSpan(npage)
	if s != nil {
		goto HaveSpan
	}
	throw("grew heap, but no adequate free span found")

HaveSpan:
	// Mark span in use.
	if s.state != mSpanFree {
		throw("candidate mspan for allocation is not free")
	}
	if s.npages < npage {
		throw("candidate mspan for allocation is too small")
	}

	// First, subtract any memory that was released back to
	// the OS from s. We will re-scavenge the trimmed section
	// if necessary.
	memstats.heap_released -= uint64(s.released())

	if s.npages > npage {
		// Trim extra and put it back in the heap.
		t := (*mspan)(h.spanalloc.alloc())
		t.init(s.base()+npage<<_PageShift, s.npages-npage)
		s.npages = npage
		h.setSpan(t.base()-1, s)
		h.setSpan(t.base(), t)
		h.setSpan(t.base()+t.npages*pageSize-1, t)
		t.needzero = s.needzero
		// If s was scavenged, then t may be scavenged.
		start, end := t.physPageBounds()
		if s.scavenged && start < end {
			memstats.heap_released += uint64(end - start)
			t.scavenged = true
		}
		s.state = mSpanManual // prevent coalescing with s
		t.state = mSpanManual
		h.freeSpanLocked(t, false, false, s.unusedsince)
		s.state = mSpanFree
	}
	// "Unscavenge" s only AFTER splitting so that
	// we only sysUsed whatever we actually need.
	if s.scavenged {
		// sysUsed all the pages that are actually available
		// in the span. Note that we don't need to decrement
		// heap_released since we already did so earlier.
		sysUsed(unsafe.Pointer(s.base()), s.npages<<_PageShift)
		s.scavenged = false
	}
	s.unusedsince = 0

	h.setSpans(s.base(), npage, s)

	*stat += uint64(npage << _PageShift)
	memstats.heap_idle -= uint64(npage << _PageShift)

	//println("spanalloc", hex(s.start<<_PageShift))
	if s.inList() {
		throw("still in list")
	}
	return s
}
```

### mheap.pickFreeSpan

```go
func (h *mheap) pickFreeSpan(npage uintptr) *mspan {
	tf := h.free.find(npage)
	ts := h.scav.find(npage)

	// Check for whichever treap gave us the smaller, non-nil result.
	// Note that we want the _smaller_ free span, i.e. the free span
	// closer in size to the amount we requested (npage).
	var s *mspan
	if tf != nil && (ts == nil || tf.spanKey.npages <= ts.spanKey.npages) {
		s = tf.spanKey
		h.free.removeNode(tf)
	} else if ts != nil && (tf == nil || tf.spanKey.npages > ts.spanKey.npages) {
		s = ts.spanKey
		h.scav.removeNode(ts)
	}
	return s
}
```

### mheap.grow

```go
func (h *mheap) grow(npage uintptr) bool {
	ask := npage << _PageShift
	v, size := h.sysAlloc(ask)
	if v == nil {
		print("runtime: out of memory: cannot allocate ", ask, "-byte block (", memstats.heap_sys, " in use)\n")
		return false
	}

	// Scavenge some pages out of the free treap to make up for
	// the virtual memory space we just allocated. We prefer to
	// scavenge the largest spans first since the cost of scavenging
	// is proportional to the number of sysUnused() calls rather than
	// the number of pages released, so we make fewer of those calls
	// with larger spans.
	h.scavengeLargest(size)

	// Create a fake "in use" span and free it, so that the
	// right coalescing happens.
	s := (*mspan)(h.spanalloc.alloc())
	s.init(uintptr(v), size/pageSize)
	h.setSpans(s.base(), s.npages, s)
	atomic.Store(&s.sweepgen, h.sweepgen)
	s.state = mSpanInUse
	h.pagesInUse += uint64(s.npages)
	h.freeSpanLocked(s, false, true, 0)
	return true
}
```


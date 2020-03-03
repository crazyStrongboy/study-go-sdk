### 基本结构

```go
type maptype struct {
	typ        _type
	key        *_type // 键的类型
	elem       *_type // 值的类型
	bucket     *_type // hash桶的内部类型
	keysize    uint8  // key的尺寸
	valuesize  uint8  // value的尺寸
	bucketsize uint16 // bucket的尺寸
	flags      uint32
}
```

```go
// A header for a Go map.
type hmap struct {
	count     int // map中元素的个数
	flags     uint8
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
	hash0     uint32 // hash种子

	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

	extra *mapextra // optional fields
}
```

![empty](http://images.hcyhj.cn/blogimages/map/bucket.png)



### 创建

```go
func makemap(t *maptype, hint int, h *hmap) *hmap {
	mem, overflow := math.MulUintptr(uintptr(hint), t.bucket.size)
	if overflow || mem > maxAlloc {
		hint = 0
	}

	// initialize Hmap
	if h == nil {
		h = new(hmap)
	}
	h.hash0 = fastrand()

	// 找到一个合适的B的值让其能够满足hint对应数量的数据
	// For hint < 0 overLoadFactor returns false since hint < bucketCnt.
	B := uint8(0)
	for overLoadFactor(hint, B) {
		B++
	}
	h.B = B

	
	if h.B != 0 {
		var nextOverflow *bmap
		h.buckets, nextOverflow = makeBucketArray(t, h.B, nil)
		if nextOverflow != nil {
			h.extra = new(mapextra)
			h.extra.nextOverflow = nextOverflow
		}
	}

	return h
}
```



### 新增数据

```go
func mapassign(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer {
	if h == nil {
		panic(plainError("assignment to entry in nil map"))
	}
	
	alg := t.key.alg
	hash := alg.hash(key, uintptr(h.hash0))

	// 设置一个标记，防止并发写，后面会检测这个标记
	h.flags ^= hashWriting

	if h.buckets == nil {
		h.buckets = newobject(t.bucket) // newarray(t.bucket, 1)
	}

again:
    // 获取分配到的桶的位置
	bucket := hash & bucketMask(h.B)

    // 获取选中桶对应的bmap
	b := (*bmap)(unsafe.Pointer(uintptr(h.buckets) + bucket*uintptr(t.bucketsize)))
    // 获取该key的高位hash值
	top := tophash(hash)

	var inserti *uint8
	var insertk unsafe.Pointer
	var val unsafe.Pointer
bucketloop:
	for {
		for i := uintptr(0); i < bucketCnt; i++ {
			if b.tophash[i] != top {
				if isEmpty(b.tophash[i]) && inserti == nil {
                    // 记录tophash中最早一次出现empty的情况，这里不一定是emptyRest，有可能是
                    // 删除数据时的emptyOne状态
					inserti = &b.tophash[i]
					insertk = add(unsafe.Pointer(b), dataOffset+i*uintptr(t.keysize))
					val = add(unsafe.Pointer(b), dataOffset+bucketCnt*uintptr(t.keysize)+i*uintptr(t.valuesize))
				}
				if b.tophash[i] == emptyRest {
					break bucketloop
				}
				continue
			}
			k := add(unsafe.Pointer(b), dataOffset+i*uintptr(t.keysize))
			if t.indirectkey() {
				k = *((*unsafe.Pointer)(k))
			}
            // 计算当前k与传进来的key是否相等，有可能只是高位相等，则需要继续循环
			if !alg.equal(key, k) {
				continue
			}
			// key在这个map中，需要更新
			if t.needkeyupdate() {
				typedmemmove(t.key, k, key)
			}
			val = add(unsafe.Pointer(b), dataOffset+bucketCnt*uintptr(t.keysize)+i*uintptr(t.valuesize))
			goto done
		}
        // 当前桶已经没有存储数据的位置了，则需要从overflow溢出桶中查询位置
		ovf := b.overflow(t)
		if ovf == nil {
			break
		}
		b = ovf
	}
    // 扩容触发点
    if !h.growing() && (overLoadFactor(h.count+1, h.B) || tooManyOverflowBuckets(h.noverflow, h.B)) {
		hashGrow(t, h)
		goto again // Growing the table invalidates everything, so try again
	}

	if inserti == nil {
		// 所有的当前桶都是满的，需要构建一个新的溢出桶
		newb := h.newoverflow(t, b)
		inserti = &newb.tophash[0]
		insertk = add(unsafe.Pointer(newb), dataOffset)
		val = add(insertk, bucketCnt*uintptr(t.keysize))
	}

	// store new key/value at insert position
	if t.indirectkey() {
		kmem := newobject(t.key)
		*(*unsafe.Pointer)(insertk) = kmem
		insertk = kmem
	}
	if t.indirectvalue() {
		vmem := newobject(t.elem)
		*(*unsafe.Pointer)(val) = vmem
	}
	typedmemmove(t.key, insertk, key)
	*inserti = top
    // 自增map中元素的个数
	h.count++

done:
	if h.flags&hashWriting == 0 {
        // 并发的写，标记在上面通过h.flags ^= hashWriting设置
		throw("concurrent map writes")
	}
    // 取消并发写的标记
	h.flags &^= hashWriting
	if t.indirectvalue() {
		val = *((*unsafe.Pointer)(val))
	}
    // 返回对应值的地址
	return val
}

```

![empty](http://images.hcyhj.cn/blogimages/map/map-bucket.png)

结合上面这张图，咱们举个简单的例子，假设hash(key)之后定位到了1号桶，然后高8位的值为top，那么咱们就会遍历1号桶对应的tophash去查询对应的地址。用简单的伪代码表现为：

```go
loop:
    for i:=0;i<8;i++{
        if tophash[i] == top{
           return &bucket+tophash_size+8*key_size+i*value_size  
        }
        if tophash[i] == empty{
            return &bucket+tophash_size+8*key_size+i*value_size  
        }
        //前面都没找到，则要去预留桶中去找位置,继续循环预留桶中的tophash
        tophash = overflowBucket.tophash
        goto loop
    }
```



### 获取数据

```go
func mapaccess1(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer {
	// map为空或者元素数量为0时快速返回
	if h == nil || h.count == 0 {
		if t.hashMightPanic() {
			t.key.alg.hash(key, 0) // see issue 23734
		}
		return unsafe.Pointer(&zeroVal[0])
	}
    // 标记并发修改位
	if h.flags&hashWriting != 0 {
		throw("concurrent map read and map write")
	}
	alg := t.key.alg
	hash := alg.hash(key, uintptr(h.hash0))
	m := bucketMask(h.B)
	b := (*bmap)(add(h.buckets, (hash&m)*uintptr(t.bucketsize)))
	// 获取hash高8位的值
	top := tophash(hash)
bucketloop:
	for ; b != nil; b = b.overflow(t) {
		for i := uintptr(0); i < bucketCnt; i++ {
			if b.tophash[i] != top {
                // 提前返回，emptyRest表示在它之后不会存在非空值
				if b.tophash[i] == emptyRest {
					break bucketloop
				}
				continue
			}
			k := add(unsafe.Pointer(b), dataOffset+i*uintptr(t.keysize))
			if t.indirectkey() {
				k = *((*unsafe.Pointer)(k))
			}
            // 防止仅top位相同的hash冲突
			if alg.equal(key, k) {
				v := add(unsafe.Pointer(b), dataOffset+bucketCnt*uintptr(t.keysize)+i*uintptr(t.valuesize))
				if t.indirectvalue() {
					v = *((*unsafe.Pointer)(v))
				}
				return v
			}
		}
	}
	return unsafe.Pointer(&zeroVal[0])
}

```

获取数据代码看起来是最容易懂的，首先定位到桶的位置，然后遍历该桶的tophash，查询key对应的位置，当前桶查询不到则继续在其溢出桶中进行递归检索即可。



### 删除数据

```go
func mapdelete(t *maptype, h *hmap, key unsafe.Pointer) {
	if h == nil || h.count == 0 {
		if t.hashMightPanic() {
			t.key.alg.hash(key, 0) // see issue 23734
		}
		return
	}
	if h.flags&hashWriting != 0 {
		throw("concurrent map writes")
	}

	alg := t.key.alg
	hash := alg.hash(key, uintptr(h.hash0))

	// 标记此时正在执行修改操作
	h.flags ^= hashWriting
	// 按低位进行桶的选择
	bucket := hash & bucketMask(h.B)
    
    if h.growing() {
		growWork(t, h, bucket)
	}

	b := (*bmap)(add(h.buckets, bucket*uintptr(t.bucketsize)))
	bOrig := b
	top := tophash(hash)
search:
	for ; b != nil; b = b.overflow(t) {
		for i := uintptr(0); i < bucketCnt; i++ {
			if b.tophash[i] != top {
				if b.tophash[i] == emptyRest {
					break search
				}
				continue
			}
			k := add(unsafe.Pointer(b), dataOffset+i*uintptr(t.keysize))
			k2 := k
			if t.indirectkey() {
				k2 = *((*unsafe.Pointer)(k2))
			}
			if !alg.equal(key, k2) {
				continue
			}
			// Only clear key if there are pointers in it.
			if t.indirectkey() {
				*(*unsafe.Pointer)(k) = nil
			} else if t.key.kind&kindNoPointers == 0 {
				memclrHasPointers(k, t.key.size)
			}
			v := add(unsafe.Pointer(b), dataOffset+bucketCnt*uintptr(t.keysize)+i*uintptr(t.valuesize))
			if t.indirectvalue() {
				*(*unsafe.Pointer)(v) = nil
			} else if t.elem.kind&kindNoPointers == 0 {
				memclrHasPointers(v, t.elem.size)
			} else {
				memclrNoHeapPointers(v, t.elem.size)
			}
            // 标记删除位，证明这个位置已被删除
			b.tophash[i] = emptyOne
	
			if i == bucketCnt-1 {
                // 如果是该tophash的最后一个，则需要判断overflow中tophash的第一位是否是emptyRest
				if b.overflow(t) != nil && b.overflow(t).tophash[0] != emptyRest {
					goto notLast
				}
			} else {
                // 如果不是当前tophash的最后一个，则直接判断当前tophash的下一个是否是emptyRest
				if b.tophash[i+1] != emptyRest {
					goto notLast
				}
			}
             // 当前值之后不存在其他的数据，则置为emptyRest
             // 需要把该位置之前除非空之外的所有emptyOne位置也置为emptyRest
			for {
				b.tophash[i] = emptyRest
				if i == 0 {
					if b == bOrig {
						break // beginning of initial bucket, we're done.
					}
					// Find previous bucket, continue at its last entry.
					c := b
					for b = bOrig; b.overflow(t) != c; b = b.overflow(t) {
					}
					i = bucketCnt - 1
				} else {
					i--
				}
				if b.tophash[i] != emptyOne {
					break
				}
			}
		notLast:
			h.count--
			break search
		}
	}
	// 抛出并发修改异常
	if h.flags&hashWriting == 0 {
		throw("concurrent map writes")
	}
	h.flags &^= hashWriting
}

```

删除数据的逻辑不是很复杂，但是里面对tophash中的状态进行了置换，增加了阅读代码的复杂度。这里主要是emptyOne和emptyRest这两个状态的转化，里面有几个临界点，下面一一说明：

1. 当前删除元素是tophash中的最后一个元素，则需要检查它的下一个溢出桶tophash的第一个元素是否是emptyRest。如果不是，则直接return。
2. 当前删除元素不是tophash中的最后一个元素，则检查当前tophash的下一个元素是否是emptyRest即可。如果不是，则直接return。
3. 前两个条件都不满足的情况下，则需要将当前位置置为emptyRest，同时还要检测它前面的emptyOne状态的元素，将其均置为emptyRest，如下图所示：

![empty](http://images.hcyhj.cn/blogimages/map/empty.png)



### 扩容

扩容的触发点在`mapassign`这个过程中，它主要检测了两个阈值：

1. 达到了负载因子的上限，数量超过了6.5*bucket_num。这时候说明大部分的桶可能都快满了，如果插入新元素，有大概率需要挂在 overflow 的桶上。
2. 溢出桶的数量过多。例如频繁的删除操作，导致溢出桶中很多位置都是emptyOne。



```go
func hashGrow(t *maptype, h *hmap) {
	bigger := uint8(1)
	if !overLoadFactor(h.count+1, h.B) {
        // 这里表示是溢出桶数量数量过多，但是map中元素数量还达不到过载，则无需增长桶数量
        // 分配一个与之前一样的即可，将所有数据摊平，减少溢出桶的数量
		bigger = 0
		h.flags |= sameSizeGrow
	}
	oldbuckets := h.buckets
	newbuckets, nextOverflow := makeBucketArray(t, h.B+bigger, nil)

	flags := h.flags &^ (iterator | oldIterator)
	if h.flags&iterator != 0 {
		flags |= oldIterator
	}
	// commit the grow (atomic wrt gc)
	h.B += bigger
	h.flags = flags
	h.oldbuckets = oldbuckets
	h.buckets = newbuckets
	h.nevacuate = 0
	h.noverflow = 0

	if h.extra != nil && h.extra.overflow != nil {
		// Promote current overflow buckets to the old generation.
		if h.extra.oldoverflow != nil {
			throw("oldoverflow is not nil")
		}
		h.extra.oldoverflow = h.extra.overflow
		h.extra.overflow = nil
	}
	if nextOverflow != nil {
		if h.extra == nil {
			h.extra = new(mapextra)
		}
		h.extra.nextOverflow = nextOverflow
	}
    // 实际的拷贝并不是在此处进行，而是在growWork() and evacuate()时自动进行的。
}

```



### 总结

在增加、查询、删除这几个操作中，其实有很多共性的地方，比如都需要先定位桶的位置，然后通过其hash值的高8位去tophash中查询是否有与此值匹配的位置存在，查询不到，则去overflow溢出桶的tophash中继续查询即可。只不过在其中做了很多逻辑判断和业务上的一些优化以及扩缩容处理，这样让代码看起来你略显复杂，咱们将问题拆解，再一步步去看，便会简单许多。
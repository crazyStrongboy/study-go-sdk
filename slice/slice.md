##### 内存泄漏问题
先看一段代码：
```go
	s := []int{-1, 1, 2, 3, 4}
	s1 := s[1:3:3]
```
当我们对s进行取其子切片后不在使用s切片，虽然s[0]和s[4]都已经不会使用了，但是由于s1共享着s的底层元素，导致s[0]和s[4]在短时间内不会释放掉，直到s1不再被使用。下面验证一下这个问题：
```go
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := []int{-1, 1, 2, 3, 4}
	s1 := s[1:3:3]
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s1))
	fmt.Println(*(*int)(unsafe.Pointer(sh.Data - unsafe.Sizeof(s[0])))) // -1
}
```

运行上面的代码发现结果为-1，这个-1为s[0]对应的元素值。由此可见s1与s0共享了这一块连续的内存地址。
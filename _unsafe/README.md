# unsafe

unsafe中公有四种方法：
1. unsafe.Pointer(&x) //任何类型都可以转换成Pointer，相当于一个万能指针
2. unsafe.Sizeof(x) // 返回x类型占用的字节数
3. unsafe.Offsetof(x) //返回相对于基址的偏移量
4. unsafe.Alignof(x) //返回对齐位数

example:
```go
package main

import (
	"fmt"
	"unsafe"
)

type A struct {
	_    int
	Name string
}

func main() {
	a := A{Name: "hejun"}
	align := unsafe.Alignof(a)
	fmt.Printf("align: %d\n", align)
	size := unsafe.Sizeof(a)
	fmt.Printf("size: %d\n", size)
	offset := unsafe.Offsetof(a.Name)
	fmt.Printf("offset: %d,name: %s\n", offset, *(*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&a)) + offset)))
}
```

打印结果：
```go
align: 8
size: 24
offset: 8,name: hejun
```
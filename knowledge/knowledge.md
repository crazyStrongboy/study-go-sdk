参考于[go101](https://github.com/golang101/golang101)，[Go官方FAQ](https://golang.google.cn/doc/faq#different_method_sets)
#### Go常见问题

##### 1. 如何强制代码的使用者必须带字段名称的组合字面量来申明结构体?

结构体中包含一个不被导出的变量，一般可使用零尺寸字段。

```go
package foo

type Config struct {
	_    [0]int
	//age int
	Name string
}
```

```go
package main

import "github.com/crazyStrongboy/study-go-sdk/knowledge/foo"

func main() {
	_ = foo.Config{[0]int{}, "name"}
	// 在go version 1.12.6版本中,编译器不会报错,但是运行时会出现如下错误:
	// implicit assignment of unexported field '_' in foo.Config literal
}
```

尽量不要把零尺寸的字段放在结构体的最后一位,因为这样会增大结构体的尺寸,增大不必要的内存.

example:

```go
type A struct {
	_    [0]int
	Name string
}

type B struct {
	Name string
	_    [0]int
}

func main() {
	aSize := unsafe.Sizeof(A{})
	bSize := unsafe.Sizeof(B{})
	fmt.Printf("aSize: %d; bSize: %d\n", aSize, bSize)
	// aSize: 16; bSize: 24
}
```

由此可见,零尺寸放到最后一位,使得该结构体多用了8个字节,具体原因如下:

> 一个可寻址的结构体，它里面的所有字段都是可寻址的，如果该结构体的最后一个字段是零尺寸，那么取最后一个字段的地址将会超过分配给该结构体的内存块的地址，有可能指向了另一个内存块(别的结构体)。如果一个内存块被至少一个及以上的指针引用，它将不会被gc回收。所以一个活跃的指针存储着非零尺寸的结构体的最后一个字段越界，它将阻止gc去回收另一个内存块，导致内存泄漏。为避免这个问题，go编译器会在有需要时给最后的一个零尺寸字段进行补齐部分字节，使其不会返回超过分配给该结构体的内存块的地址。

> ps:如果一个结构体的全部字段都是零尺寸的，那么就不需要填充字节，标准编译器会专门处理零尺寸的内存块。



##### 2.为什么两个`nil`值有时候会不相等？

先看一下interface{}的定义：

```go
// emptyInterface is the header for an interface{} value.
type emptyInterface struct {
	typ  *rtype
	word unsafe.Pointer
}
```

由上面的定义可以了解到，一个接口由两个元素组成，一个是类型T（typ）,一个值V（word）。当申明一个接口类型的变量值为3，即``` var a interface{} = 3```，实际上底层定义为```T=*int，V=3```。值V被作为接口的动态值，因为在运行过程中，一个接口变量会拥有不同的V，但是T是相同的。

一个interface是nil的表现为类型为nil，值不进行设置。实际上，空接口将总是持有一个空的类型，即T=nil。如果我们存储一个空指针```*int```到一个接口中，该变量会表现为```T=*int,V=nil```,因此一个接口即使在值为空的情况下也不为空。



example:

```go
package main

import "fmt"

func main() {
	var pi *int = nil
	var pb *bool = nil
	var x interface{} = pi // T=*int V=nil
	var y interface{} = pb // T=*bool V=nil
	var z interface{} = nil

	fmt.Println(x == y)   // false
	fmt.Println(x == nil) // false
	fmt.Println(y == nil) // false
	fmt.Println(x == z)   // false
	fmt.Println(y == z)   // false
}
```

##### 3.为什么T和*T有不同的方法集？

- 一个`T`类型的值可以调用`*T`类型声明的方法，前提是这个T的值可以寻址的情况下。编译器在调用指针所属的方法前，会先自动取`T`的地址值。因为不是任何`T`值都是可以寻址的，所以并非任何T值都能调用`*T`的方法。
- 一个`*T`类型的值可以调用`T`类型声明的所有方法，因为所有接引用都是合法的。

所以说，`*T`类型值的方法集是`T`类型方法集的超集。反之不然。

> PS:即使在编译器可以在方法的调用过程中自动取地址，但是在操作不当的情况下，也会导致调用过程中的更改不会体现到源值上。

example:
```go
package main

import "fmt"

type T struct {
	Name string
}

func (t *T) Change() {
	t.Name = "change"
}

func main() {
	t := T{Name: "init"}
	fmt.Println("before: ", t.Name)//before:  init
	change(t)
	fmt.Println("after: ", t.Name)  //after:  init
}

func change(t T) {
	t.Change()
}
```
上面例子可以清楚的看到，虽然`change`方法中自动取了t的地址，然后调用属于`*T`的方法。但是并未改变源值，这是因为传进`change`方法的参数只是一个副本，取址后自然也不会是源值的地址。

##### 4.怎样保证一个类型实现了某个接口？
可以使用编译器去检查类型`T`是否实现了接口`I`，比如：
```go
type T struct{}
var _ I = T{}       // Verify that T implements I.
var _ I = (*T)(nil) // Verify that *T implements I.
```
如果`T`或者`*T`没有实现接口I，那么在编译时期就会捕获到这个错误。
>PS:如果一个接口要想只被内部使用，可以在接口中声明一个不被导出的方法即可。

##### 5.能够转换[]T到[]interface{}吗？
不能够直接的进行转换，因为两种类型在内存中的含义不一样。需要单独一个个拷贝到目的slice中。

example：
```go
    t := []int{1, 2, 3, 4}
    s := make([]interface{}, len(t))
    for i, v := range t {
        s[i] = v
    }
```

##### 6.如果T1和T2有相同的底层类型，能够强转[]T1到[]T2吗？
```go
    type T1 int
    type T2 int
    var t1 T1
    var x = T2(t1) // OK
    var st1 []T1
    var sx = ([]T2)(st1) // NOT OK 编译不通过
```

##### 7.为什么map中不允许slice作为key？
map的查询需要一个比较器，但是slice没有实现这个比较器。因为对于切片类型比较器比较难定义，这里面要涉及到很多考虑，比如：浅拷贝与深拷贝的比较，指针与值的比较，如何处理递归类型等等。

在目前GO版本中，已经为结构体和数组实现了比较器，所以这些类型能够作为map的key。slice到目前还不被支持。

下列类型不支持比较：
- 映射（map)
- 切片
- 函数
- 包含不可比较字段的结构体
- 元素类型为不可比较类型的数组类型

##### 8.什么时候函数参数通过值传递？
在Go中任何时候调用方法都是通过值传递。函数总是拷贝所要传递的参数，就好像赋值一样将值赋值给参。例如，传递一个int类型的值到函数中相当于拷贝一个int值，传递一个指针类型的参数相当于拷贝一个指针，而不是这个指针指向的数据。

##### 9.我应该将方法定义在值上还是指针上？
```go
    func (s *MyStruct) pointerMethod() { } // method on pointer
    func (s MyStruct)  valueMethod()   { } // method on value
```
这个问题其实可以将方法的接收者看成一个参数，这样就可以很清楚的定义接收者是值还是指针了。可以从下面几点进行分析：
- 方法是否需要改变接收者里面的内容，如果需要，则接收者必须是指针。
- 考虑到高效性，如果接收者是一个很大的结构体，则将接收者定义为指针会减少不必要的内存开销。
- 考虑到一致性，如果该接收者的部分方法必须要指针去接收，则可以考虑将另外的部分方法均定义为指针方法。

##### 10.将一个闭包作为goroutine运行时会发生什么？
```go
func main() {
    done := make(chan bool)

    values := []string{"a", "b", "c"}
    for _, v := range values {
        go func() {
            fmt.Println(v)
            done <- true
        }()
    }

    // wait for all goroutines to complete before exiting
    for _ = range values {
        <-done
    }
}
```
运行上述程序，也许会期待a,b,c的顺序输出，但是实际上输出的是c,c,c。这是因为每一次迭代使用了相同的变量v，所以说每一个闭包都共享同一个实例。当闭包函数运行时，`fmt.Println`将会打印v的值，但是在goroutine运行前v已经被修改，所以打印的最终都是c。大家也可以把这个理解为闭包传递的是v的地址，并不是v的拷贝。

可以通过下面两个方法去避免上述问题：
```go
    for _, v := range values {
        go func(u string) {
            fmt.Println(u)
            done <- true
        }(v) // 进行参数值的拷贝
    }
```
```go
    for _, v := range values {
        v := v // create a new 'v'，每次传递的都是一个新的v变量，地址不同
        go func() {
            fmt.Println(v)
            done <- true
        }()
    }
```

##### 11.哪些值可以取地址，哪些值不可以取地址？

以下的值是可寻址的，因此可以被取地址：
- 变量: `&x`
- 可寻址的结构体的字段: `&point.X`
- 可寻址的数组的元素：`&a[0]`
- 任意切片的元素（无论是可寻址切片或不可寻址切片）：`&s[1]`
- 指针解引用（dereference）操作：`&*x`
- composite literal类型: `&struct{ X int }{1}`

下列情况x是不可以寻址的，你不能使用&x取得指针：

- 字符串中的字节
- map对象中的元素(因为Go中map实现中元素的地址是变化的，这意味着寻址的结果是无意义的)
- 接口对象的动态值(通过type assertions获得)
- 常数
- package 级别的函数
- 方法method (用作函数值)
- 中间值(intermediate value):
    - 函数调用
    - 显式类型转换
    - 各种类型的操作 （除了指针引用pointer dereference操作 *x):
        - 通道接收操作
        - 子字符串操作
        - 子切片操作
        - 加减乘除等运算符
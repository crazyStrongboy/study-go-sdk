## channel
#### 简介

熟悉Go的人都知道，它提倡着不要通过共享内存来通讯，而要通过通讯来共享内存。Go提供了一种独特的并发同步技术来实现通过通讯来共享内存，此技术即为通道。我们可以把一个通道看作是在一个程序内部的一个FIFO数据队列。 一些协程可以向此通道发送数据，另外一些协程可以从此通道接收数据。



#### Example
介绍一下简单的用法：
```go
func main() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	t := <-c
	fmt.Println(t)
}
```
几个注意点：

1. 向一个nil通道中发送一个值，将会永久阻塞。
2. 向一个已关闭的通道中发送一个值，将会导致panic。
3. 可以从关闭的通道中读取值，缓冲区为空时，读取的是通道类型的零值。
4. 重复关闭一个通道也会导致panic。
5. 通道的元素值的传递都是复制过程，且至少被复制过一次以上。（直接复制到receiver中经过一次复制，通过缓冲区的话则经历了两次复制）



#### channel有两种类型，Unbuffered channels与Buffered channels

##### Unbuffered channels
```go
	c:=make(chan int)
```
它是一个阻塞型channel，必须要receiver也准备好的情况下，sender才能够将消息投递到c中去。可以结合下图进行思考一波：

![unbuffer_channel](images/unbuffer_channel.png)

##### Buffered channels
```go
	c:=make(chan int,1)
```
在buf未满之前，它是一个非阻塞型channel，sender可以将符合channel类型的值投递到channel中去，它内部会自己维护一个队列。当buf满了之后，sender会阻塞。可以结合下图进行思考一波：

![buffer_channel](images/buffer_channel.png)

#### 几种应用模式

##### for-range

```go
func forRange() {
	c := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		c <- i
	}
	close(c)
	for v := range c {
		fmt.Println(v)
	}
}
```

1. 在进行for-range一个通道时，该循环将源源不断的从通道中获取数据，直到此通道关闭并且它的缓冲队列中为空为止。
2. 这里的通道一定不能是单向发送通道（`chan <- int`）。
3. 当for-range一个空通道时，将会永久阻塞

##### select-case

```go
func selectCase() {
	c := make(chan int, 1)
	c <- 1
	close(c)
	select {
	case <-c:
		fmt.Println("xxxx")
	default:
		fmt.Println("aaaa")
	}
}
```

1. 每个`case`关键字后必须跟随一个通道接收数据操作或者一个通道发送数据操作。

2. 所有的非阻塞`case`操作中将有一个被随机选择执行（而不是按照从上到下的顺序），然后执行此操作对应的`case`分支代码块。

3. 在所有的`case`操作均为阻塞的情况下，如果`default`分支存在，则`default`分支代码块将得到执行； 否则，当前协程将被推入所有阻塞操作中相关的通道的发送数据协程队列或者接收数据协程队列中，并进入阻塞状态。

   

#### 源码分析
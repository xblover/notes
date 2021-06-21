package main

import (
	"fmt"
	"sync"
	"unsafe"
	"time"
)

func main() {
	fmt.Println("vim-go")
	/**
		go鼓励使用CSP模型，以通信来代替内存共享，实现并发安全
		Do't communicate by sharing memory, share memory by communicating

		作为CSP的核心，channel是显示的。 要求操作双方必须知道数据类型和具体通道，并不关心另一端操作者的身份和数量
		可如果另一端未准备妥当，或者消息未能及时处理，会阻塞。

		从底层实现上来说，通道只是一个队列。同步模式下，发送和接收双方配对然后直接复制数据给对方，如果配对失败，则
		置入等待队列，直到另一方出现才被唤醒。异步模式抢夺的则是数据缓冲槽，发送方要求有空槽可供写入，而接收方则要求
		有缓冲数据可读。需求不符时，同样加入等待队列，直到有另一方写入数据或腾出空槽后被唤醒。
		

	**/
	//hi()
	//test1()
	//test2()

	//	okidom()
	//forrange()

	ready()


	/**
		向已关闭通道发送数据，引发panic
		从已关闭接收数据，返回已缓冲数据或零值.

	**/
}


func ready() {

	var wg sync.WaitGroup
	ready := make(chan struct{})

	for i:=0; i<3 ; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			println(id, ":ready.")     //运动员装备就绪
			<-ready
			println(id, ":running...") //等待发令
		}(i)
	}

	time.Sleep(time.Second)
	println("Ready? Go!")

	close(ready)

	wg.Wait()

}

func hi() {

	// 除传递消息（数据）外，通道还常被用作事件通知

	done := make(chan struct{}) // 结束事件
	c := make(chan string)      // 数据传输通道

	go func() {
		s := <-c // 接收消息
		println(s)
		close(done) //关闭通道，作为结束通知
	}()

	c <- "hi!" // 发送消息
	<-done     // 阻塞，直到有数据或管道关闭

}

func test1() {

	//缓冲区大小仅仅是内部属性，不属于类型组成部分。另外通道变量本身就是指针，可用相等等操作符判断
	// 是否为同一对象或nil

	var a, b chan int = make(chan int, 3), make(chan int)
	var c chan bool

	println(a == b)   // 通道变量本身就是指针，可用相等等操作符判断是否为同一对象
	println(c == nil) // 通道变量本身就是指针，可用相等等操作符判断是否为nil

	fmt.Printf("%p, %d, %d\n", a, unsafe.Sizeof(a), unsafe.Sizeof(b)) //缓冲区大小仅仅是内部属性，不属于类型组成部分
}

func test2() {
	a, b := make(chan int), make(chan int, 3)

	b <- 1
	b <- 2

	// 内置函数len和cap返回缓冲区大小和当前已缓存数量；对于同步通道则都返回0，
	// 据此可判断通道是同步还是异步的。
	println("a:", len(a), cap(a)) //
	println("b:", len(b), cap(b)) // 内置函数len和cap返回缓冲区大小和当前已缓存数量；

}

func okidom() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done) //确保发出结束通知

		for {
			x, ok := <-c
			if !ok { // 据此判断通道是否被关闭
				return
			}

			println(x)
		}
	}()

	c <- 1
	c <- 2
	c <- 3
	close(c)

	<-done
}

func forrange() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)

		for x := range c { //循环获取消息，直到通道被关闭
			println(x)
		}
	}()

	c <- 1
	c <- 2
	c <- 3
	close(c)

	<-done
}

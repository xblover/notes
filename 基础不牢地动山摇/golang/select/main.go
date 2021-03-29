package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("vim-go")
	test1()



}

func test1() {
	// 如果要处理多个通道，可选用select语句，。它会随机选择一个可用通道做收发操作
	
	var wg sync.WaitGroup
	wg.Add(2)

	a, b := make(chan int), make(chan int)

	go func() {    //接收端
		defer wg.Done()

		for {
			var (
				name string
				x int
				ok bool
			)
		
			select { //随机选择可用的channel 接收数据
			case x, ok = <-a:
				name = "a"
			case x, ok = <-b:
				name = "b"
			}

			if !ok { // 如果任一通道关闭，则终止接收
				return
			}

			println(name, x) //输出接收的数据信息
		}
	}()

	go func() {   // 发送端
		defer wg.Done()
		defer close(a)
		defer close(b)

		for i := 0; i<10; i++ {
			select {   //随机选择发送channel
			case a <- i:
			case b <- i * 10:
			}
		}


	}()

	wg.Wait()

}

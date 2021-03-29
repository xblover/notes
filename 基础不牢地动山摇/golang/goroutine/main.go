package main

import (
	"fmt"
	"time"
	"sync"
) 
/**
	并发(concurrency)与并行(parallesim)
	并发：逻辑上具备同时处理多个任务的能力
	并行：物理上在同一时刻执行多个并发任务


	关键字go并非执行并发操作，而是创建一个并发任务单元。新建任务被放置在系统队列中，等待调度器安排合适系统线程去获取执行权
	每个任务单元除保存函数指针，调用函数外，还会分配执行所需的栈内存空间。相比系统默认MB级别的线程栈， goroutine自定义栈仅需
	2kb, 所以创建成千上万的并发任务。自定义栈采取按需分配策略，在需要扩容时扩容，最大能到GB规模


	你看到的和产出的一切都在以并发方式运行。垃圾回收，系统监控，网络通信，文件读写，还有用户并发任务等，这些都需要一个高效
	且聪明的调度器来指挥协调。

	1.语句go func() 创建G ，
	2.放入P本地队列（或平衡到全局队列）
	3.唤醒或新建M执行任务
	4. 进入调度循环schedule
	5. 竭力获取待执行G 任务并执行
	6，清理现场，重新进入调度循环





	因为G初始栈仅有2KB，且创建操作只是在用户空间简单非配对象，远比进入内核态分配线程要简单的多，调度器让多个M进入调度循环，不停
	获取执行任务，所以我们才能创建成千上万个并发任务。



**/
func main() {
	fmt.Println("vim-go")
	//	wait()

	//wait()

//	waitGroup()
//	waitGroup1()

	waitGroup2()

	
	
}

func waitGroup2() {

	// 在多处使用wait阻塞，它们都能接收到通知

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		wg.Wait()   //等待归零，解除阻塞
		println("wait exit.")
	}()

	go func() {
		time.Sleep(time.Second)
		println("done.")
		wg.Done()
	}()             // 递减计数

	wg.Wait()       // 等待归零，解除阻塞
	println("main exit.")

}
/**
func waitGroup1() {
	var wg sync.WaitGroup

	// 尽管waitgroup.add 实现了原子操作，但建议在goroutine外累加计数器，以免Add尚未执行，wait已经退出
	go func() {
		wg.Add(1)    //来不及设置
		println("hi!")
	}

	wg.Wait()
	println("exit.")
}
**/
func waitGroup() {
	//要等待多个任务，推荐使用sync.WaitGroup. 通过设定计数器，让每个goroutine,在退出前递减，直至归零时解除阻塞

	var wg sync.WaitGroup

	for i := 0; i < 10 ; i++ {
		wg.Add(1) //累加计数

		go func(id int) {
			defer wg.Done()  //递减计数

			time.Sleep(time.Second)
			println("goroutine", id, "Done.")
		}(i)
	}

	println("main ...")
	wg.Wait()   // 阻塞，直到计数归零
	println("main exit.")
}

func wait() {

	//进程退出时不会等待并发任务结束，可用通道channel 阻塞，然后发出退出信号
	
	exit := make(chan struct{}) //创建通道，因为仅是通知，数据无实际意义

	go func() {
		time.Sleep(time.Second)
		println("goroutine done.")

		close(exit) // 关闭通道，发出信号
	}()

	println("main ...") 
	<-exit				// 如通道关闭，立即解除阻塞
	println("main exit.")

}

func wait1() {
	//exit := make(chan struct{}) //创建通道，因为仅是通知，数据无实际意义
	
	println("--------------------------------------------------------\n")
	go func() {
		time.Sleep(time.Second)
		println("goroutine done.")

	//		close() // 关闭通道，发出信号
	}()

	println("main ...") 
					// 如通道关闭，立即解除阻塞
	println("main exit.")

}









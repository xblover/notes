package main

import "fmt"

func consumer(data chan int, done chan bool) {
	for x := range data {
		println("recv:",x) // 接收数据直到通道被关闭
	}

	done <- true   // 通知main, 消费结束
}

func producer(data chan int) {
	for i:=0; i < 4; i++ {
		data <- i     //发送数据
	}

	close(data)
}


func main() {
	fmt.Println("vim-go")
	done := make(chan bool) //用于接收消费结束信号
	data := make(chan int)  //数据管道

	go consumer(data,done) //启动消费者
	go producer(data)      //启动生产者

	<-done                 // 阻塞，直到消费者发回结束符号
}

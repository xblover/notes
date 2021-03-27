package main

import "fmt"

type RWMutex struct {
	w Mutex     //复用互斥锁
	writeSem uint32   //写等待读
	readerSem uint32  //读等待写
	readerCount int32 //存储了当前正在执行的读操作数量
	readerWait int32  //表示当写操作被阻塞时等待的读操作个数
}

func main() {
	fmt.Println("vim-go")
}

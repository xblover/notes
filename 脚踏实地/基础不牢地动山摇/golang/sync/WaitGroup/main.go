package main

import (
	"fmt"
	"sync"
)

/**
	WaitGroup 可以解决一个goroutine等待多个goroutine同时结束的场景，比较常用的场景如：
	多线程下载，后端worker启动了多个消费者干活
**/

func main() {
	fmt.Println("vim-go")
	var wg sync.WaitGroup
	for i :=0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			worker(i)
		}(i)
	}
	wg.Wait()
}

func worker(i int) {
	fmt.Println("worker: ", i)
}


/*
	
*/

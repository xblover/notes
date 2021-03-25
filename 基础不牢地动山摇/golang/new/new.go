package main

import "fmt"

/*
	new函数也可为引用类型分配内存，但这是不完整创建。以map为例，它仅仅分配了字典类型本身（实际就是一个指针包装）所需内存，并没有分配
	键值存储内存，也没有初始化散列桶等内部属性，因此无法正常工作
*/
func main() {
	fmt.Println("vim-go")
	p:= new(map[string]int) // 函数new 返回指针
	m := *p
	m["a"] = 1
	fmt.Println(m)  // panic: assignment to entry in nil map 
}

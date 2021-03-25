package main

import "fmt"
/**
	所谓引用类型特指slice, map, channel 这三种预定义类型
	
	相比数字，数组等类型，引用类型拥有更复杂的数据结构。除分配内存外，还须初始化一系列属性，如指针，长度，甚至包括哈希分布，数据对列等

	内置函数new 按指定类型长度分配零值内存，返回指针，不会关心类型内部构造和初始化方式。而引用类型必须用make函数创建，编译器会将make转换
	为目标类型专用的创建函数（或指令），以确保完成全部内存分配和相关属性的初始化
**/
func mkslice() []int {
	s := make([]int, 0, 10)
	s = append(s, 100)
	return s
}

func mkmap() map[string]int {
	m := make(map[string]int)
	m["a"] = 1
	return m
}

func main() {
	fmt.Println("vim-go")
	m := mkmap()
	println(m["a"])

	s := mkslice()
	println(s[0])
}

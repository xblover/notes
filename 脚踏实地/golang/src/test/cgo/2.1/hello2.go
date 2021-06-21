package main

/*
#include <stdio.h>

static void SayHello(const char* s) {
	puts(s);
}
*/
import "C"

func main() {
	//fmt.Println("vim-go")
	// 定义一个SayHello 的C函数来实现打印，
	// 然后从GO环境中调用
	C.SayHello(C.CString("hello string\n"))
}

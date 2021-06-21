package main

//void SayHello(const char* s);
import "C"

func main() {
	//fmt.Println("vim-go")
	C.SayHello(C.CString("hello world\n"))
}

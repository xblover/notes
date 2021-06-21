package main

//#include <stdio.h>
import "C"

func main() {
	//fmt.Println("vim-go")
	C.puts(C.CString("Hello world\n"))
}

package main

import "fmt"
/**
	零长度对象的地址是否相等和具体的实现版本有关，但肯定不等于nil
	即使长度为0，可该对象依然是合法存在的，拥有合法内存地址，这与nil语义完全不同
**/
func main() {
	fmt.Println("vim-go")
	var a, b struct{}

	println(&a, &b)
	println(&a == &b, &a == nil)

}

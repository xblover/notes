#  数组

## 概述
	数组是由相同类型元素的集合组成的数据结构，计算机会为数组分配一块连续的内存来保存其中的元素。我们可以利用数组中元素的索引快速访问特定元素。

	数组作为一种基本数据类型，我们会从数组中存储的元素类型和数组最大能存储的元素个数两个维度描述数组。

	go语言数组在初始化之后大小就无法改变，存储元素类型相同，但大小不同的数组类型在go语言看来也是完全不同的，只有两个条件都相同才是同一类型。

	```go
	func NewArray(elem *Type, bound int64) *Type {
		if bound < 0 {
			Fatalf("NewArray: invalid bound %v", bound)
		}

		t := New(TARRAY)
		t.Extra = &Array{Elem: elem, Bound: bound}
		t.SetNotInHeap(elem.NotInHeap())
		return t
	}
	```
	编译期间的数组类型是由上述的cmd/compile/internal/types.NewArray 函数生成，该类型包含两个字段，分别是元素类型Elem和数组大小Bound, 这两个字段共同构成了数组类型。


## 初始化

	go语言的数组有两种不同的创建方式，
	```go
		arr1 := [3]int{1,2,3}    //显示的指定数组大小
		arr2 := [...]int{1,2,3}  //使用 [...]T 声明数组
	```
	上述两种声明方式在运行期间得到的结果是完全相同的，后一种声明方式在编译期间会转换成前一种，这也就是编译器对数组大小的推导。


	// todo: 上限推导，语句转换 细节分析


## 访问和赋值

	无论是在栈上还是在静态存储区，数组在内存中都是一连串的内存空间，我们通过数组开头的指针，元素的数量以及元素类型占的空间大小表示数组。如果
我们不知道数组中元素的数量，访问时可能发生越界；而如果不知道数组中元素类型的大小，就没有办法知道应该一次取多少字节的数据，无论丢失了那个信息，
我们都无法知道这片连续的内存空间到底存储了什么数据。

	//todo: 数组访问越界分析

## 总结
	
	文章大致总结了go语言数组的定义与结构（元素类型，元素个数），数组的初始化方式（显示方式，需要编译器推导大小的方式），
	还有数组的访问与赋值方式

##  扩展阅读

 - [ 数组 ](https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-array/#31-%E6%95%B0%E7%BB%84)
 - [Arrays, slices (and strings): The mechanics of 'append'](https://blog.golang.org/slices)
 - [Array vs Slice: accessing speed](https://stackoverflow.com/questions/30525184/array-vs-slice-accessing-speed)





	











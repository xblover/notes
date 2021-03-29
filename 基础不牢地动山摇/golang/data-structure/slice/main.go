package main

import (
	"fmt"
	"unsafe"
	"reflect"
)

func main() {
	
	test()
	// 扩容
	testExpansion()

}

func testExpansion() {
	slice := []int{10,20,30,40}
	newSlice := append(slice, 50)
	fmt.Printf("Before slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("Before newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	newSlice[1] += 10
	fmt.Printf("After slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("After newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))

	//output:
	// Before slice = [10 20 30 40], Pointer = 0xc0000ae030, len = 4, cap = 4
	// Before newSlice = [10 20 30 40 50], Pointer = 0xc0000ae048, len = 5, cap = 8
	// After slice = [10 20 30 40], Pointer = 0xc0000ae030, len = 4, cap = 4
	// After newSlice = [10 30 30 40 50], Pointer = 0xc0000ae048, len = 5, cap = 8

	// 新的切片更改了一个值，并没有影响到原来的数组，新切片指向的数组是一个全新的数组。并且cap容量也发生了变化。
	// go中切片扩容的策略是这样的：
	// 首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap）,最终容量（newcap）就是新申请的容量（cap）
	// 否则判断，如果旧切片的长度小于1024，则最终容量（newcap）就是旧容量（old.cap）的两倍，即（newcap=doublecap）
	// 否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量(old.cap)开始循环增加原来的25%，...
	// 如果最终容量计算值溢出，则最终容量就是新申请的容量。


	array := [4]int{10, 20, 30, 40}
	slice1 := array[0:2]
	newSlice1 := append(slice1, 50)
	fmt.Printf("--------------------------------------------------------------------------------------------------\n")   
	fmt.Printf("Before slice = %v, Pointer = %p, len = %d, cap = %d\n", slice1, &slice1, len(slice1), cap(slice1))
	fmt.Printf("Before newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice1, &newSlice1, len(newSlice1), cap(newSlice1))
	newSlice1[1] += 10
	fmt.Printf("After slice = %v, Pointer = %p, len = %d, cap = %d\n", slice1, &slice1, len(slice1), cap(slice1))
	fmt.Printf("After newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice1, &newSlice1, len(newSlice1), cap(newSlice1))
	fmt.Printf("After array = %v\n", array)
	// todo :新数组，老数组？ 解释





	
		
}

func test() {
	fmt.Println("vim-go")

	// 从slice中得到一块内存地址
	s := make([]byte, 200)
	ptr := unsafe.Pointer(&s[0])
	fmt.Printf("%v\n", ptr)
	
	// 在go的反射中就存在一个与之对应的数据结构SliceHeader,我们可以用它来构造一个slice
	var o []byte
	var length int
	var ptr1 unsafe.Pointer
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&o)))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(ptr1)
	fmt.Printf("%v\n",sliceHeader)
	
}

type slice struct {
	array unsafe.Pointer
	len int
	cap int
}

// slicecopy is used to copy from a string or slice of pointerless elements into a slice
// slicecopy用于将字符串或无指针元素的切片复制到切片中
func slicecopy(toPtr unsafe.Pointer, toLen int, fromPtr unsafe.Pointer, fromLen int, width uintptr) int {
	// 如果源切片或者目标切片有一个长度为0，那么就不需要拷贝，直接return
	if fromLen == 0 || toLen == 0 {
		return 0
	}

	// n 记录下源切片或者目标切片较短的那一个的长度
	n := fromLen
	if toLen < n {
		n = toLen
	}

	// 如果入参width = 0, 也不需要拷贝了，返回较短的切片的长度
	if width == 0 {
		return n
	}
	
	size := uintptr(n) * width
	// 如果开启了竞争检测
	if raceenabled {
		callerpc := getcallerpc()
		pc := funcPC(slicecopy)
		racereadrangepc(fromPtr, size, callerpc, pc)
		racewriterangepc(toPtr, size, callerpc, pc)
	}

	// 如果开启了 the memory sanitizer (msan)
	if msanenabled {
		msanread(fromPtr, size)
		msanwrite(toPtr, size)
	}

	if size == 1 { // common case worth about 2x to do here 普通案例的价值约为此处的2倍
		// TODO: is this still worth it with new memmove impl?
		// 

		//如果只有一个元素，那么指针直接转换即可
		*(*byte)(toPtr) = *(*byte)(fromPtr) // known to be a byte pointer 已知是字节指针

	} else {
		// 如果不止一个元素，那么就把size 个bytes 从fromPtr copy 到 toPtr 
		memmove(toPtr, fromPtr, size)
	}

	return n
}



/**
func growslice(et *_type, old slice, cap int) slice {
	if raceenabled {
		callerpc := getcallerpc()
		racereadrangepc(old.array, uintptr(old.len*int(et.size)), callerpc, funcPC(growslice))
	}
	if msanenabled {
		msanread(old.array, uintptr(old.len*int(et.size)))
	}

	// 如果新要扩容的容量比原来的容量还要小，这代表要缩容了，那么可以直接报panic了。
	if cap < old.cap {
		panic(errorString("growslice: cap out of range"))
	}

	if et.size == 0 {
		
		// 如果当前切片的大小为0，还调用了扩容方法，那么就新生成一个新的容量的切片返回。
		return slice{unsafe.Pointer(&zerobase), old.len, cap} 
	}


	// 这里就是扩容的策略
	newcap := old.cap
	doublecap := newcap + newcap
	if cap > doublecap {
		newcap = cap
	} else {
		if old.cap < 1024 {
			newcap = doublecap
		} else {
			// Check 0 < newcap to detect overflow and prevent an infinite loop。
			// 检查0 < newcap 以检测溢出并防止无限循环。
			for 0 < newcap && newcap < cap {
				newcap += newcap / 4
			}

			// Set newcap to the requested cap when the newcap calculation overflowed
			// 当newcap计算溢出时，将newcap设置为请求的cap 
			if newcap <= 0 {
				newcap = cap
			}
		}
	}

	// 计算新的切片的容量，长度。
	var overflow bool
	var lenmem, newlenmem, capmem uintptr
	// Specialize for common values of et.size.
	// For 1 we don't need any division/multiplication
	// For sys.PtrSize, compiler will optimize division/multiplication into a shift by a constant.
	// For powers of 2, use a variable shift.
	// 专门用于et.size的通用值。
	// 对于1，我们不需要任何除法/乘法
	// 对于sys.PtrSize, 编译器会将除法/乘法优化为常数。
	// 对于2的幂，请使用可变移位。
	switch {
	case et.size == 1:
		lenmem = uintptr(old.len)
		newlenmem = uintptr(cap)
		capmem = roundupsize(uintptr(newcap))
		overflow = uintptr(newcap) > maxAlloc
		newcap = int(capmem)
	case et.size == sys.PtrSize:
		lenmem = uintptr(old.len) * sys.PtrSize
		newlenmem = uintptr(cap) * sys.PtrSize
		capmem = roundupsize(uintptr(newcap) * sys.PtrSize)
		overflow = uintptr(newcap) > maxAlloc/sys.PtrSize
		newcap = int(capmem / sys.PtrSize)
	case isPowerOfTwo(et.size):
		var shift uintptr
		if sys.PtrSize == 8 {
			// Mask shift for better code generation
			// 掩码移位可更好地生成代码
			shift = uintptr(sys.Ctz64(uint64(et.size))) & 63
		} else {
			shift = uintptr(sys.Ctz32(uint32(et.size))) & 31
		}
		lenmem = uintptr(old.len) << shift
		newlenmem = uintptr(cap) << shift
		capmem = roundupsize(uintptr(newcap) << shift)
		overflow = uintptr(newcap) > (maxAlloc >> shift)
		newcap = int(capmem >> shift)
	default:
		lenmem = uintptr(old.len) * et.size
		newlenmem = uintptr(cap) * et.size
		capmem, overflow = math.MulUintptr(et.size, uintptr(newcap))
		capmem = roundupsize(capmem)
		newcap = int(capmem / et.size)
	}

	// The check of overflow in addition to capmem > maxAlloc is needed 
	// to prevent an overflow which can be used to trigger a segfault
	// on 32bit architectures with this example program:
	//
	// type T [1<<27 + 1]int64
	//
	// var d T
	// var s []T
	//
	// func main() {
	// 	  s = append(s, d,d,d,d)
	//    print(len(s),"\n")
	// }
	//
	// 除了 capmem > maxAlloc 外，还需要检查溢出，以防止溢出，此示例程序可用于触发32位体系结构上的段错误：
	// 
	if overflow || capmem > maxAlloc {
		panic(errorString("growslice: cap out of range"))
	}

	var p unsafe.Pointer
	if et.ptrdata == 0 {
		// 在老切片后面继续扩充容量
		p = mallocgc(capmem, nil, false)
		// The append() that calls growslice is going to overwrite from old.len to cap (which will be the new length).
		// only clear the part that will not be overwritten
		// 调用growslice 的append() 将从old.len 覆盖为cap（这将是新的长度）。
		// 仅清除不会被覆盖的部分。

		// 先将P地址加上新的容量得到新切片容量的地址，然后将新切片容量地址后面的 capmem-newlenmem
		memclrNoHeapPointers(add(p, newlenmem), capmem-newlenmem)
	} else {
		// Note: can't use rawmem (which avoids zeroing of memory), because then GC can scan uninitialized memory.
		// 注意：不能使用rawmem(这样可以避免内存清零)，因为GC可以扫描未初始化的内存。

		//重新申请新的数组给新切片
		// 重新申请capmen 这个大的内存地址，并且初始化为0值
		p = mallocgc(capmem, et, true)
		if lenmem > 0 && writeBarrier.enabled {
			// only shade the pointers in old.array since we know the destination slice p
			// only contains nil pointers because it has been cleared during alloc.
			// 由于我们知道目标切片p仅包含nil指针，因为在分配过程中已将其清除，因此仅在old.array 中隐藏了指针。
			bulkBarrierPreWriteSrcOnly(uintptr(p), uintptr(old.array), lenmem-et.size+et.ptrdata)
		}
	}

	//将 lenmem这个多个 bytes从old.array地址 拷贝到p的地址处
	memmove(p, old.array, lenmem)
	// 返回最终新切片，容量更新为最新扩容之后的容量
	return slice{p, old.len, newcap}
}
**/





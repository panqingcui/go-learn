package main

import "fmt"

//如果一个指针变量存指向的又是另一个指针变量的指针地址，则称这个指针变量为指向指针的指针变量
func main() {
	var a = 300
	var ptr *int
	//指针ptr 内存地址
	ptr = &a
	//指向指针ptr 内存地址
	var pptr **int
	pptr = &ptr

	fmt.Printf("变量a=%d \n", a)
	fmt.Printf("指针变量ptr = %d,%x \n", *ptr, ptr)

	fmt.Printf("指向指针的指针变量ptr = %d,%x,%x \n", **pptr, *pptr, pptr)
}

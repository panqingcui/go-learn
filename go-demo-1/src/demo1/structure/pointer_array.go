package main

import "fmt"

//常量
const MAX = 3

func main() {
	//正整型数组
	a := []int{1, 2, 3}
	//指针数组
	var ptr [MAX]*int
	//整数地址赋给指针数组
	var i int
	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i]
	}
	for i = 0; i < MAX; i++ {
		fmt.Printf("\n 输出 指针对应的地址 %x \n", ptr[i])
	}
	for i = 0; i < MAX; i++ {
		fmt.Printf("\n 输出 指针对应的值 %d \n", *ptr[i])
	}
}

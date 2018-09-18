package main

import "fmt"

func main() {
	//
	var a int
	fmt.Println("a 变量的内存", a)
	fmt.Printf("a 变量的内存 %d \n", a)
	fmt.Printf("&a 变量的地址 %v \n", &a)
	fmt.Println("&a 变量的地址", &a)
	//保存某个变量的地址，需要指针类型 *int 保持int的地址 **int 保持 *int的地址
	var p *int
	//指针变量指向谁，把a的变量地址赋给 p
	p = &a
	fmt.Printf("p = %v ,&a = %v \n", p, &a)
	*p = 666 //*p 操作的不是p的内存，而是p所指向的内存
	fmt.Printf("p = %v ,a = %v \n", p, a)
	// new 分配内存
	var n *int
	n = new(int)
	*n = 123
	fmt.Printf("无效指向n = %v \n", n)

}

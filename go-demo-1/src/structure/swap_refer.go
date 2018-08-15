package main

import "fmt"

/**
引用传递，实际上传递值在内存中的实际地址，传递指针地址
*/
func main() {
	var a, b int = 100, 200
	fmt.Printf("交换前的值 a is %d,b is %d", a, b)
	//&a 指向a的指针，a的内存地址
	//&b 指向b的指针，b的内存地址
	swap(&a, &b)

	fmt.Printf("交换后的值 a is %d,b is %d", a, b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x //保存 x的 内存地址
	*x = *y   //保存y 的内存地址
	*y = temp //保存原x的内存地址
}

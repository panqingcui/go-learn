package main

import "fmt"

//一个指针变量可以指向任何一个值得内存地址，它指向那个值得内存地址
func main() {
	var a int = 10
	//声明指针变量
	var b *int
	//指针变量的存储地址
	b = &a
	fmt.Printf("\n a 的变量地址:%x", &a)
	fmt.Printf("\n b 变量存储的指针地址：%x ", b)
	fmt.Printf("\n * b 指针值 %d \n", *b)
	//空指针nil 与其他语言的null,nil,none含义一致，都指代零值或空值
	var ptr *int
	fmt.Printf("\n空指针 %x", ptr)
	fmt.Printf("\n 测试 %d\n", 11111)
	//空指针的判断
	if ptr != nil {
		fmt.Printf("\n 空指针 %x", ptr)
	}

	if ptr == nil {
		fmt.Printf("\n 空指针 %x \n", ptr)
	}

}

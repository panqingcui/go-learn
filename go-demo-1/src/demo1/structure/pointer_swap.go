package main

import "fmt"

func main() {
	//	fmt.Println("vim-go")
	//声明两个
	var a int = 100
	var b int = 200
	fmt.Printf("交换前 a=%d \n", a)
	fmt.Printf("交换前 b=%d \n", b)
	swap(&a, &b)
	fmt.Printf("交换后 a=%d \n", a)
	fmt.Printf("交换后 b=%d \n", b)

}

func swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp

}

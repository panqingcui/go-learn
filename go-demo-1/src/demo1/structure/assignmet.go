package main

import "fmt"

func main() {
	//命令变量的赋值
	var x = 1
	fmt.Println("x", x)
	y := 1
	fmt.Println("y", y)
	//通过指针间接赋值
	//var *p = true
	//结构体变量赋值
	//var person.name = "bob"
	//
	//var count[x] = count[x] * scale

	fmt.Println("15,20得最大公约数：", gcd(15, 20))
	fmt.Println("斐波那契数列：5", fib(5))
}
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

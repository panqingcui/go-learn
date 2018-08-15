package main

import "fmt"

func swap(x, y int) int {
	var temp int //局部变量
	temp = x     //临时变量 ，保存x的值
	x = y        // y 值赋给x
	y = temp     // temp -> y
	return temp
}

func main() {
	var a, b int = 100, 200
	fmt.Printf("交换前 a %d,b %d", a, b)

	swap(a, b)

	fmt.Printf("值交换后 a%d,b%d", a, b)

}

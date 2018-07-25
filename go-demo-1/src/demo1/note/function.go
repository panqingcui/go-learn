package main

import "fmt"

//多返回值，a,b 参数调换顺序
func vals(a, b int) (int, int) {
	return b, a
}
func sum(a, b int) int {
	return a + b
}

//命名返回值
func split(sum int) (x, y int) {
	x = sum - 4
	y = sum + 4
	return
}
func main() {
	a, b := vals(1, 2)
	_, c := vals(1, 2)
	fmt.Println("a: ", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	//求和函数
	sum := sum(1, 2)
	fmt.Println("sum:", sum)
	fmt.Println(split(1))
}

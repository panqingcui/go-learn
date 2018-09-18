package main

import "fmt"

func test3(x int) {
	fmt.Println("----", 100/x)
}
func main() {
	//defer 延迟调用。函数退出时调用
	//defer 先进的后出
	//函数异常的时候,defer 一样执行
	defer fmt.Println("bbbbbbb------")
	defer fmt.Println("ccccccc------")
	defer fmt.Println("aaaaaaaa")
	//内存异常
	//defer test3(0)
	defer fmt.Println("ddddddd------")
	a, b := 10, 20

	defer func(x int) {
		fmt.Println("defer 闭包 ", x, b)
	}(a) //(a)

	a = 100
	b = 200
	fmt.Println("a,b", a, b)

}

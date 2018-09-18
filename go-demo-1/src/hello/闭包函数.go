package main

import "fmt"

func test1() int {
	//函数呗调用时，分配空间
	var x int
	x++
	return x //调用完毕，释放空间
}

//返回值--匿名函数 func() --匿名函数的类型 int--返回一个函数类型
func test2() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	fmt.Println("普通包测试", test1())
	fmt.Println("普通包测试", test1())
	//返回值，匿名函数，变量接收
	f := test2()
	fmt.Println("闭包测试", f()) //1
	fmt.Println("闭包测试", f()) //4

}

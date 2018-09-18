package main

import "fmt"

//func add(args ...int)  {
//	fmt.Println("不定参数 len ",len(args))
//	fmt.Println("不定参数 ",args)
//	for i:=0;i<len(args) ;i++  {
//		fmt.Println("不定参数 ",args[i])
//	}
//	for i:= range args {
//		fmt.Println("不定参数 ",args[i])
//	}
//
//	for _,data := range args {
//		fmt.Println("不定参数 ",data)
//	}
//	//全部参数传递
//	fmt.Println("全部参数---")
//	test(args...)
//	fmt.Println("截至到 2，不包括 2---")
//	//从2开始传递
//	test(args[:2]...)
//	fmt.Println("从2开始，包括2本身---")
//	test(args[2:]...)
//}
/**
1到100
*/
func addNumber() (sum int) {
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return sum
}
func addNumber1(i int) int {
	if i == 1 {
		return 1
	}
	return i + addNumber1(i-1)
}
func addNumber2(i int) int {
	if i == 100 {
		return 100
	}
	return i + addNumber2(i+1)
}
func test(args ...int) {
	for _, data := range args {
		fmt.Println("test ----->>>>>>>不定参数 ", data)

	}
}
func add(a, b int) int {
	return a + b
}
func minus(a, b int) int {
	return a - b
}

type f func(int, int) int

func Call(a, b int, fTest f) (result int) {
	fmt.Println("---->>>>>Call")
	result = fTest(a, b)
	return
}
func main() {
	//add(1,2,3,4,5)
	var sum int
	//1到100
	sum = addNumber()
	fmt.Println("test ----->>>>>>>1--100求和 ", sum)
	fmt.Println("递归算法 1--100求和")
	fmt.Println("---------", addNumber1(100))
	fmt.Println("---------", addNumber2(1))
	fmt.Println("函数类型---是一种基本类型")
	var t f
	t = add
	fmt.Println("1+2 ", t(1, 2))
	t = minus
	fmt.Println("1+2 ", t(1, 2))
	fmt.Println("函数参数接收 type 类型")
	fmt.Println("----add", Call(1, 2, add))
	fmt.Println("----minus", Call(1, 2, minus))
	fmt.Println("------匿名函数--通过括号函数调用")
	func(a, b int) {
		fmt.Println("匿名函数", a+b)
	}(1, 2)
	fmt.Println("------匿名函数--通过赋值接收")
	f := func(a, b int) {
		fmt.Println("匿名函数", a+b)
	}
	f(1, 2)

	fmt.Println("------匿名函数--通过赋值接收,返回值")
	max, min := func(a, b int) (max, min int) {
		if a > b {
			max = a
			min = b
		} else {
			max = b
			min = a
		}
		return
	}(1, 2)
	fmt.Println("mix,min -", max, min)

}

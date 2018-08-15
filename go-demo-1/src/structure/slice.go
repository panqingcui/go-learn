package main

import "fmt"

func main() {
	//go 数组，声明及赋值，遍历
	var a [10]int
	var i, j int
	for i = 0; i < 10; i++ {
		a[i] = i + 10
	}
	for j = 0; j < 10; j++ {
		fmt.Printf("a[%d] = %d \n", j, a[j])
	}
	//数组的初始化
	var b = [5]int{1, 2, 3}

	for j = 0; j < 5; j++ {
		fmt.Printf("--j %d \n", b[j])
	}

	//动态的数组
	//go 语言切片是数组的抽象
	//go语言中的数组长度不可变

	//declare
	//	var identifier []type
	//	var slice1 []type = make([]type,len)
	//	slice2 := make([]type,len)
	//init
	//直接初始化切片，[]表示切片类型{1，2，3}初始化值，其cap=len=3
	//	s :=[] int{1,2,3}
	//数组arr的引用
	//	s := arr[:]
	// lem(),cap()
	var numbers = make([]int, 3, 5)
	printSlice(numbers)

}

func printSlice(x []int) {
	fmt.Printf("x len %d,cap %d,slice %v \n", len(x), cap(x), x)
}

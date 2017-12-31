package main

import "fmt"

func main() {
	//GO 中变量规则
	// var 变量名称 类型 = 表达式
	var i = 0
	fmt.Printf("%d \n", i)
	var num int = 30
	fmt.Print("\n num:", num)
	var sum = num + 10
	fmt.Printf("\n num = %d sum = %d:", num, sum)
	var f float32 = 20.0
	y := 1
	fmt.Printf("\n f number is %f", f)
	fmt.Printf("\n  f type %T", f)
	fmt.Printf("\n y", y)

}

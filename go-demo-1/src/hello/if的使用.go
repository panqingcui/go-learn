package main

import "fmt"

func main() {

	s := "富豪"
	if s == "富豪" { //左括号与if 同一行
		fmt.Println("不用上班了。。。。")
	} else {
		fmt.Println("继续上班。。。。")
	}
	if a := 10; a == 10 {
		fmt.Println("同一语句。")
	}
	var num int
	fmt.Println("请输入参数")
	fmt.Scan(&num)
	switch num {
	case 1:
		fmt.Println("1 ", num)
	case 2:
		fmt.Println("2 ", num)
	default:
		fmt.Println("3", num)
	}
}

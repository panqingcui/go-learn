package main

import "fmt"

func main() {
	var num1 int = 100
	var num2 int = 200
	var result int
	result = max(num1, num2)
	fmt.Printf("%d %d 最大值%d", num1, num2, result)
}

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

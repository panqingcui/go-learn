package main

import "math"
import "fmt"

/**
创建一个内置函数，并直接使用
*/
func main() {
	sqrt := func(x float64) float64 {
		return math.Ceil(x)
		//	return math.Sqrt(x)
	}
	fmt.Println(sqrt(3))
	fmt.Println(sqrt(3.1))
}

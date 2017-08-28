package main

import "fmt"

/**
* 正整型数组，常量，及遍历
 */
//数组长度 ,常量声明 不需要 var
const MAX = 3

func main() {
	a := []int{1, 2, 3}
	for b := 0; b < MAX; b++ {
		fmt.Printf(" a[%d]==%d \n", b, a[b])
	}
}

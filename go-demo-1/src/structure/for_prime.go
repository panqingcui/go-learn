package main

import "fmt"

func main() {
	//输出2至100的素数
	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j < i/j; j++ {
			fmt.Printf("\ni  %d,%d", i, i/j)
			if i%j == 0 {
				break //存在因数
			}
		}
		if j > i/j {
			fmt.Printf("\n%d 是素数", i)
		}
	}
}

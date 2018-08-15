package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
重复的
*/
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	//fmt.Printf("\n%d",counts)
	for input.Scan() {
		counts[input.Text()]++
		fmt.Printf(input.Text())
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

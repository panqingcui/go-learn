package main

import "fmt"

func main() {
	var a int
	for a = 0; a < 20; a++ {
		if a == 10 {
			break
		}
		fmt.Println(a)
	}
}

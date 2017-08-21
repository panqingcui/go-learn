package main

import "fmt"

func main() {

	var a [10]int
	for i := 0; i < 10; i++ {
		a[i] = i + 100
	}
	for j := 0; j < 10; j++ {
		fmt.Println(j, a[j])
	}
	fmt.Println("vim-go")
	var balance = [5]float32{1.0, 2.0, 3.0, 4.0, 5.0}
	var x int
	for x = 0; x < 5; x++ {
		fmt.Println("x:", balance[x])
	}
}

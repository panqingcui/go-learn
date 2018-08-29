package main

import (
	"../swap"
	"fmt"
)

/**
 */
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Printf("max %d \n", max(1, 2))
	fmt.Printf("max %d \n", max(1, 2))
	a := "f"   
	b := "s"

	x, y := swap.Swap(a, b)
	fmt.Printf("swap %s,%s", x, y)
	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	for i := 0; i < len(balance); i++ {
		fmt.Printf(" balance[%d]|%f \n", i, balance[i])
	}
}

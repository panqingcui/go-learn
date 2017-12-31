package main

import "fmt"

const boilingF = 212.0

func main() {
	var i = boilingF
	var c = (i - 32) * 5 / 9
	fmt.Printf("%gF or %g", i, c)

}

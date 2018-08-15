package main

import "fmt"

const (
	a = iota
	b = 1 << iota
	c
	d
)

func main() {
	fmt.Println("", a, b, c, d)
}

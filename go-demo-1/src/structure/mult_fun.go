package main

import "fmt"

func main() {
	a, b := swap("BeiJing is the captal of ", "China")

	fmt.Printf("%s,%s", a, b)
}

func swap(x, y string) (string, string) {

	return x, y
}

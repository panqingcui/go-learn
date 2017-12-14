package main

import "fmt"

func main() {
	fmt.Println(count(10))
	stacked()
	c := Car{model: "this is line number 8"}
	defer c.PrintModel()
	c.model = "this is line number 10"

}

type Car struct {
	model string
}

func (c Car) PrintModel() {
	fmt.Println(c.model)

}
func count(i int) (n int) {

	defer func(i int) {
		fmt.Println(">>>>>", n, i)
		n = n + i
	}(i)

	i = i * 2
	n = i
	return
}

func stacked() {
	defer func() {
		fmt.Println("last")

	}()

	defer func() {
		fmt.Println("first")
	}()

}

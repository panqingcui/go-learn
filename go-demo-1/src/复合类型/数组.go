package main

import "fmt"

func main() {
	var a [10]int
	//var b [5] int
	for i := 0; i < len(a); i++ {
		a[i] = i + 1
	}

	for i, data := range a {
		fmt.Printf("a[%d]  = %d \n", i, data)
	}
	var b [3]int = [3]int{1, 2, 3}
	for i, data := range b {
		fmt.Printf("b[%d]  = %d \n", i, data)
	}
	//指定元素赋值
	c := [5]int{2: 10, 3: 15}
	for i, data := range c {
		fmt.Printf("指定元素----->c[%d]  = %d \n", i, data)
	}
	d := [5]int{1, 2, 3, 4, 5}
	e := [5]int{1, 2, 3, 4, 5}
	fmt.Println("d==e", d == e)
	var f [5]int
	f = d
	fmt.Println("f=", f)
}

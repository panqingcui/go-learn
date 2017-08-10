package main

import "fmt"
import "unsafe"

const ()

func main() {

	const WIDTH = 5.1
	const LENGTH = 3
	const a, b, c = 1, true, "test"
	const area = WIDTH * LENGTH
	fmt.Printf("面积 area %f", area)
	println()
	println(a, b, c)
	const d = "abc"
	const e = len(d)
	const f = unsafe.Sizeof(e)
	println(d, e, f)
}

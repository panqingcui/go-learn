package main

func main() {
	//var a int
	numbers := [6]int{1, 2, 3, 4}

	for a := 0; a < 10; a++ {
		println("a=", a)
	}
	for i, x := range numbers {
		println("numbers 值", i, x)
	}
}

package main

import "fmt"

func main() {

	for a := 0; a < 10; a++ {
		if a == 8 {
			fmt.Printf("跳过%d \n", a)
			continue
		}
		fmt.Println(a)
	}

}

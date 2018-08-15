package main

import "fmt"

func main() {
	var marks = 90
	var grade = "B"
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	}
	switch {
	case grade == "A":
		fmt.Println("优秀")
	}
	fmt.Println("vim-go")
}

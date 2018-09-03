package main

import "fmt"

func main() {
	var a byte
	a = 'a'
	fmt.Printf("a = %c \n",a)
	fmt.Println("a = ",a)

	fmt.Printf("a = %v \n",a)
	var t int
	t = int(a)
	fmt.Printf("printf c t = %c \n",t)
	fmt.Printf("printf v t = %v \n",t)
	fmt.Println("println t = ",t)


}

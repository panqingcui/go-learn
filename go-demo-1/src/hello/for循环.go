package main

import "fmt"

func main() {
	//for 初始化;条件判断;条件变化{}
	sum := 0
	for i := 1; i <= 10000; i++ {
		sum += i
	}
	fmt.Println("sum = ", sum)
	str := "abc"
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d],%c \n", i, str[i])
	}
	//迭代元素，一个元素的位置，一个元素本身
	for i, data := range str {
		fmt.Printf("str[%d],%c \n", i, data)
	}

	for i := range str {
		fmt.Printf("str[%d],%c \n", i, str[i])
	}

	for i, _ := range str {
		fmt.Printf("str[%d],%c \n", i, str[i])
	}
}

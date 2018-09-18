package main

import (
	"fmt"
	"os"
)

func main() {
	//接收用户传递的参数，都是以字符的方式传递
	list := os.Args
	n := len(list)
	fmt.Println("----len ", n)
	for i := 0; i < len(list); i++ {
		fmt.Println("----遍历输入参数 ", i, list[i])
	}

	for i, data := range list {
		fmt.Println("----遍历输入参数 ", i, data)

	}
}

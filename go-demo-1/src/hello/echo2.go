package main

import (
	"fmt"
	"os"
)

/**
  _ 占位符
*/
func main() {
	s, sep := "", "1 "
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println("s:=", s)
}

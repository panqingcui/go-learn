package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
	log.Println("init")
	log.Println("%v", os.Stdout)
}
func main() {
	fmt.Println("vim-go")
}

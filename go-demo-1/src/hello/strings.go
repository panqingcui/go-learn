package main

import "fmt"
import (
	"os"
	"strings"
)

/*
  strings 包的使用
*/

func main() {
	fmt.Printf("s")
	fmt.Printf("\n" + os.Args[0])
	//strings.Join 中可以将字符，根据字符连接
	s := []string{"foo", "bar", "baz"}
	fmt.Println("\n strings join:" + strings.Join(s, "."))
	//大写
	s1 := "xiaomi"
	fmt.Println("\n stirngs upper:", strings.ToUpper(s1), strings.ToUpper(s1))

}

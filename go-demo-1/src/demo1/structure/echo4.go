package main

import (
	"flag"
	"fmt"
	"strings"
)

//flag 命令解析
// 返回 指针
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", "str", "separator")

func main() {
	//flag.Var(&flagVal, "name", "help message for flagname")
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	fmt.Println("flag args->>", flag.Args())
	fmt.Println("sep has value", *sep)
	fmt.Println("ip has value ", *n)
	fmt.Printf("ip has value %", n)
	//fmt.Println("flagvar has value ", flagvar)
	if !*n {
		fmt.Println()
	}

	x := 1
	//如果指针名字为p，那么可以说“p指针指向变量x”，或者说“p指针保存了x变量的内存地址”
	//同时*p表达式对应p指针指向的变量的值。
	p := &x
	fmt.Println("p 的指针", *p)
	*p = 2         // equivalent to x = 2
	fmt.Println(x) // "2"

	var d, y int
	fmt.Println(&d == &d, &d == &y, &d == nil) // "true false false"
	fmt.Println("f()", f() == f())

	//
	// var flagVar int
	// flag.IntVar(&flagVar, "flagname", 1234, "flag var name")
	// fmt.Println("flag var", *flagVar)
	//通过 flag.Var 方法 绑定变量，直接输出
	var flagvar int
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
	fmt.Println("flag var", flagvar)
}

func f() *int {
	v := 1
	return &v
}

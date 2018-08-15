package main

/**
展开 常量除 “只读” 以外，和变量究竟有什么不同
*/
var x = 0x100

const y = 0x200

func main() {
	println(&x, x)
	println(y)
}

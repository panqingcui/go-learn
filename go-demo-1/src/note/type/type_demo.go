package main

import (
	"fmt"
	"strconv"
)

/**
 *	类型相关测试
 */
func main() {
	var x int
	fmt.Printf("int 初始化 [%d] \n", x)
	var y = false
	fmt.Printf("boolean 初始化[%v] \n", y)

	var (
		d, z int
		_, s = 100, 100
	)
	fmt.Printf("[%v],[%v],[%v] \n", d, z, s)

	b := 10
	fmt.Printf("[%v]\n", &b)
	//b,c := 20,30
	//fmt.Printf("[%v] [%v] [%v]\n",b,c,&b)

	var c int
	for i := 0; i <= 10; i++ {
		c++
	}
	fmt.Printf("c:[%v] \n", c)
	//"_" 占位符 通常作为忽略占位符使用，可作表达式左值，无法获取内容,空标识符可用临时规避编译器对未使用变量和倒入包的错误检查，但请注意，踏实预置成员，没不能重新定义
	zw, _ := strconv.Atoi("12")
	fmt.Printf("%v", zw)
	// -------------常量

}

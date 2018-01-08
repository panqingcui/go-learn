package main

import (
	"unsafe"
	"fmt"
)

/**
   常量
 */
func main() {
	//const x , y int =123,0x22
	const s = "hello, world"
	const c = '我'
	const (
		i, f =1,0.123 //int ,float64(默认)
	//	b = false

	)
	{
		const x = 456
		println(x)
	}
	//println(x)
	//println(y)
	//println(b)
	println(f)
	const (
		x,y int = 99,99
		b byte =byte(x) // x 被指定为 int 类型，须显示转换为byte类型
		n = uint8(y)
		// 常量值也可以是某些编译器能计算出结果的表达式
		prtSize = unsafe.Sizeof(uintptr(1))
		strSize = len("hello,world")
		ab uint16 =120
		ac   //与上一行 ab 类型相同
		as = "abc"
	)
	fmt.Printf("%T,%v \n",ac,ac)
	println(prtSize)
	println(strSize)

	//枚举
	const (
		ax = iota //
		ay
		az
	)
	println(ax)
	println(ay)
	println(az)
	const (
		_ = iota //0
		//左移 1 * （2 的10次方）
		KB = 1<< (10*iota)
		MB
		GB
	)
	println(KB)
	println(MB)
	println(GB)
	//如中断iota 自增，则必须现实恢复，且后续自增值按行序递增
	const (
		ba = iota
		bb
		bc = 100
		bd
		be = iota
		bf
	)
	println(ba)
	println(bb)
	println(bc)
	println(bd)
	println(be)
	println(bf)

	const (
		ca = iota // int
		cb float32 =iota
		cc = iota
	)
	println(ca)
	println(cb)
	println(cc)
	//在实际编码中，建议用自定义类型实现用途明确的枚举类型。但这并不能将取值范围限定在预定义的枚举值内。
	test(red)
	test(100)
	//dx := 2
	//test(dx)
	//______-------------------------------展开

}
func test(c color)  {
	println(c)

}
type color byte //自定义类型
const (
	black color = iota //指定常量类型
	red
	blue
)

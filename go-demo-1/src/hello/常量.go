package main

import (
	"fmt"
)

func main() {
	const  a int = 10
	const (
		b = 20
		c = 30
	)
	fmt.Printf("a = %d,b = %d,c = %d \n",a,b,c)

	const a3,b3,c3 = 1,2,1.25
	fmt.Printf("a3 = %d,b3 = %d,c3 默认6位小数 = %f \n",a3,b3,c3)
	fmt.Printf("a3 = %d,b3 = %d,c3 1位小数 = %.1f \n",a3,b3,c3)

	// iota 自动+1 遇到const 自动清零
	const (
		a1 = iota
		b1 = iota
		c1 = iota
	)
	const d = iota
	fmt.Printf("a1 = %d,b1 = %d,c1 = %d ,d = %d\n",a1,b1,c1,d)
	//可以只写一个iota
	const (
		a2 = iota
		b2
		c2
	)
	fmt.Printf("a2 = %d,b2 = %d,c2 = %d \n",a1,b1,c1)

}

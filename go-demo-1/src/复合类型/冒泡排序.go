package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var a [5]int
	a = [5]int{2, 1, 3, 4, 5}
	fmt.Println("排序前", a)
	//冒泡排序,大的值交换
	for i := 0; i < len(a)-1; i++ {
		for j := i; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println("排序后", a)
	//随机生成
	var b [10]int
	//赋值
	rand.Seed(12)
	for i := 0; i < len(b); i++ {
		b[i] = rand.Intn(20)
	}
	fmt.Printf("交换前的---%v \n", b)
	//降序
	for i := 0; i < len(b)-1; i++ {
		for j := 0; j < len(b)-1-i; j++ {
			if b[j] < b[j+1] {
				b[j], b[j+1] = b[j+1], b[j]
			}
		}
	}
	fmt.Printf("交换后---->%v", b)

}

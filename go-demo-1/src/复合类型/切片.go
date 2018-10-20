package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//切片(slice) 本身并非动态数组或数组指针，他内部通过阵阵引用底层数组，设定相关属性将数据读写操作限定在指定区域内
	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//x[low:high:max] len = high - low; cap = max -low
	fmt.Printf("x[:] = %v, len = %d,cap = %d \n", x[:], len(x[:]), cap(x[:]))
	expression := x[2:5]
	fmt.Printf("x[2:5] = %v,len = %d,cap = %d \n", expression, len(expression), cap(expression))
	expression = x[2:5:7]
	fmt.Printf("x[2:5:7] = %v,len = %d,cap = %d \n", expression, len(expression), cap(expression))
	expression = x[4:]
	fmt.Printf("x[4:] = %v,len = %d,cap = %d \n", expression, len(expression), cap(expression))
	expression = x[:4]
	fmt.Printf("x[:4] = %v,len = %d,cap = %d \n", expression, len(expression), cap(expression))
	expression = x[:4:6]
	fmt.Printf("x[:4:6] = %v,len = %d,cap = %d \n", expression, len(expression), cap(expression))
	//切片初始化
	y := []int{}
	fmt.Printf("[]int{} = %v,len = %d,cap = %d \n", y, len(y), cap(y))
	y = append(y, 11)
	fmt.Printf("[]int{} = %v,len = %d,cap = %d \n", y, len(y), cap(y))
	//数组 切片的区别。数组的长度，容量是一致的，且不可以改变
	//make(slice,len,cap)
	z := make([]int, 50, 50)
	fmt.Printf("[]int{} = %v,len = %d,cap = %d \n", z, len(z), cap(z))
	//append 时 超过 容量时，容量自行扩张 *2
	z = append(z, 100)
	fmt.Printf("[]int{} = %v,len = %d,cap = %d \n", z, len(z), cap(z))
	z = append(z, 100)
	fmt.Printf("[]int{} = %v,len = %d,cap = %d \n", z, len(z), cap(z))
	fmt.Printf("z[0] = %v", z[0])
	//调整切片的值后，对元数组是有影响的。可以认为是指针指向调整
	expression[0] = 12
	fmt.Printf("原数组---> %v", x)
	//测试切片的扩容特点 .len = 0,cap = 1的切片
	a := make([]int, 0, 1)
	o := cap(a)
	for i := 0; i < 100; i++ {
		a = append(a, i)
		if n := cap(a); n > o {
			fmt.Printf("o = %d,n = %d \n", o, n)
			o = n
		}

	}
	fmt.Printf("%v", "切片copy方法测试---------------- \n")
	src := []int{1, 2, 3}
	dst := make([]int, 4, 4)
	fmt.Printf("before copy src = %v, dst = %v \n", src, dst)

	r := copy(dst, src)
	fmt.Printf("after copy src = %v, dst = %v,copy count r = %d \n", src, dst, r)
	fmt.Printf("%v", "切片当作函数使用，引用传递\n")
	s := make([]int, 20, 20)
	InitData(s)
	fmt.Printf("排序前的数据----%v \n", s)
	BubbleSort(s)
	fmt.Printf("排序后的数据----%v \n", s)

}

/**
init slice
*/
func InitData(s []int) {
	//种子
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		s[i] = rand.Intn(100)
	}
}

func BubbleSort(s []int) {
	//升序
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

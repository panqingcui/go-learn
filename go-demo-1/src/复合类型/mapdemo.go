package main

import (
	"fmt"
)

func main() {
	var m map[string]int
	fmt.Printf("m len %d \n", len(m))
	fmt.Printf("m len %v \n", m)
	//m["a"] = 1
	//不能对nil 类型进行写操作，但可以进行读操作
	fmt.Printf("m[|a] %d \n", m["a"])
	m = make(map[string]int)
	fmt.Printf("m len %d \n", len(m))
	fmt.Printf("make m len %d \n", m)
	m1 := make(map[string]int)
	m1["a"] = 1
	m1["b"] = 2
	fmt.Printf("map %v \n", m1)
	//匿名结构体
	m2 := map[int]struct {
		x int
	}{
		//可省略的key,value 类型标签
		1: {},
		2: {200},
	}
	fmt.Printf("m1---%v,m2---%v \n", m1, m2)
	fmt.Printf("m2 1:%v \n", m2[1])
	//map 的基本操作
	m3 := map[string]int{
		"a": 1,
		"b": 2,
	}
	m3["d"] = 30 //新增
	//访问不存在的键值，默认返回零值，不会引发错误，但推荐使用ok-idiom模式，
	if v, o := m3["d"]; o {
		fmt.Printf("m3['d']=%d \n", v)
	}
	//字典遍历
	for i := 0; i < 8; i++ {
		m[string('a'+i)] = i
	}
	for i := 0; i < 4; i++ {
		for k, v := range m {
			fmt.Printf("%v : %d ", k, v)
		}
		fmt.Printf("\n")
	}
	//not addressable

	type user struct {
		name string
		age  int
	}
	m4 := map[int]user{
		1: {"Tom", 18},
	}
	//对age +1
	u := m4[1]
	u.age++
	m4[1] = u
	fmt.Printf("m4 age :%v \n", m4[1])
	//通过指针方式
	m5 := map[int]*user{
		1: {"Tom", 18},
	}
	m5[1].age++
	fmt.Printf("m5 age :%v \n", m5[1])
	//
	m6 := make(map[int]int)
	for i := 0; i < 10; i++ {
		m6[i] = i
	}

	for k := range m6 {
		if k == 5 {
			m6[100] = 1000
		}
		delete(m6, k)
		fmt.Println(k, m6)
	}
}

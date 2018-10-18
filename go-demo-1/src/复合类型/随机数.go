package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//设置种子
	rand.Seed(112)
	rand.Seed(time.Now().UnixNano())
	//如果种子参数一致，产生随机数的一致
	for i := 0; i < 5; i++ {
		fmt.Println("随机数 = ", rand.Int())
	}
	for i := 0; i < 5; i++ {
		fmt.Println("随机数100内 = ", rand.Intn(100))
	}
}

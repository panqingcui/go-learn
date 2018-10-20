package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
猜数字游戏  根据输入的数字，提示用户输入的与随机数关系。
1、创建随机数
2、根据位数拆分数组
3、在键盘获取数字
*/

func main() {
	//随机数
	var randNum int
	CreateNum(&randNum)
	randSlice := make([]int, 4, 4)
	GetNum(randSlice, randNum)
	OnGame(randSlice)
}
func OnGame(randSlice []int) {
	var keyNum int
	keySlice := make([]int, 4)
	for {
		fmt.Printf("请输入数字 \n")
		fmt.Scan(&keyNum)
		if keyNum > 999 && keyNum < 10000 {
			//拆分数组
			GetNum(keySlice, keyNum)
			n := 0
			for i := 0; i < 4; i++ {
				if keySlice[i] > randSlice[i] {
					fmt.Printf("第%d位大了 \n", i+1)
				} else if keySlice[i] < randSlice[i] {
					fmt.Printf("第%d位小了 \n", i+1)
				} else {
					fmt.Printf("第%d位正确 \n123", i+1)
					n++
				}
			}
			if n == 4 {
				fmt.Printf("随机数 %v", randSlice)

				break
			}
		} else {
			fmt.Printf("输入的数字不规范")

		}
	}

}
func GetNum(randSlice []int, num int) {
	randSlice[0] = num / 1000
	randSlice[1] = (num % 1000) / 100
	randSlice[2] = (num % 100) / 10
	randSlice[3] = num % 10
}

/**
创建随机数
*/
func CreateNum(p *int) {
	//随机种子
	rand.Seed(time.Now().UnixNano())
	//死循环
	for {
		*p = rand.Intn(10000)
		if *p > 1000 {
			//保证4位数
			break
		}
	}

}

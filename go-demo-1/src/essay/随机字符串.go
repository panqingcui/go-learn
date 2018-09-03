package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
设置时间种子,否则随机数不发生变化
*/
func init() {
	//纳秒
	rand.Seed(time.Now().UnixNano())
}

// rune --> int32
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

/**
随机字符串
*/
func RandStringsRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//常量
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringsBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

/**
余数方式
*/
func RandStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
func main() {
	l := 5
	fmt.Printf("常用方式：rune,输出长度:%d,随机字符：%v \n", l, RandStringsRunes(5))
	fmt.Printf("bytes方式，输出长度:%d,随机字符：%v \n", l, RandStringsBytes(5))
	fmt.Printf("余数方式，输出长度:%d,随机字符：%v \n", l, RandStringBytesRmndr(5))

}

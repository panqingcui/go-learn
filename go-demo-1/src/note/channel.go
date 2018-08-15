package main

import (
	"fmt"
	"time"
)

// chan
var completed chan int = make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}
	//存消息
	completed <- 0
}
func main() {
	// go 启动一个 go routine
	go loop()
	//loop()
	fmt.Println(time.Now())
	//time.Sleep(time.Second)
	fmt.Println("vim-go")
	//取消息，取到消息后退出 main 阻塞
	<-completed
}

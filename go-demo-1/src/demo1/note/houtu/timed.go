package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("This is time:", time.Now())
	fmt.Println("timestamp", time.Now().Unix())
	fmt.Println("format", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("format", time.Now().Format("2006-01-03 15:04:05"))
	str_time := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	fmt.Println("timestamp->string", str_time)
	the_time, err := time.parse("2006-01-02 15:04:05", str_time)
	if err == nil {
		unix_time := the_time.Unix()
		fmt.Println("str->timestamp", unix_time)
	}

}

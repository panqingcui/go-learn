package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("当前时间字符串 time.Now() =", time.Now())
	fmt.Println("当前时间秒 time.Now().Unix() =", time.Now().Unix())
	fmt.Println("当前时间纳秒 time.Now().UnixNano() = ", time.Now().UnixNano())
	fmt.Println("当前时间毫秒 time.Now().UnixNano()/1e6 = ", time.Now().UnixNano()/1e6)
	fmt.Println("当前时间秒 time.Now().UnixNano/1e9 = ", time.Now().UnixNano()/1e9)
	fmt.Println("当前时间格式化 time.Now.Format('2006-01-02 15:04:05')", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("format", time.Now().Format("2006-01-03 15:04:05"))
	str_time := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	fmt.Println("timestamp->string", str_time)
	the_time, err := time.Parse("2006-01-02 15:04:05", str_time)
	if err == nil {
		unix_time := the_time.Unix()
		fmt.Println("str->timestamp", unix_time)
	}
	fmt.Println("当前月中的第几日 ", time.Now().Day())
	fmt.Println("当前年中的第几日 ", time.Now().YearDay())
	fmt.Println("一周中的周几 ", time.Now().Weekday())
	seconds := 10
	fmt.Println("时间段 10s ", time.Duration(seconds)*time.Second)
	fmt.Println("当前时间添加 12 小时", time.Now().Add(time.Duration(12)*time.Hour))
	t1 := time.Now()
	time.Sleep(time.Second)
	fmt.Println("两个时间点相减 time.Sub", time.Now().Sub(t1))
	fmt.Println("time location", time.Now().Location())
	fmt.Println("time Second:", time.Now().Second())
	fmt.Println("time string", time.Now().String())
	name, offset := time.Now().Zone()
	fmt.Println(" 时区信息 name,offset = ", name, " ", offset)
	//time round
	t := time.Now()
	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		time.Minute,
		time.Hour,
	}
	for _, d := range round {
		fmt.Printf("t.Round(%s) = %s \n", d, t.Round(d))
	}
}

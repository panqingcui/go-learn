package main

import (
)

/**
   1 到 100
 */
func main()  {
	//t := time.Now()
	//fmt.Println(t)
	//timestamp := t.UnixNano() / 1000000 //UnitNano获取的是纳秒，除以1000000获取毫秒级的时间戳
	//
	//for i:=2;i<200000;i++{
	//	var b bool=true
	//	for j:=2;j<i ;j++  {
	//		if i%j==0 {
	//			b=false
	//			break
	//		}
	//
	//	}
	//	if b {
	//		println(i)
	//	}
	//
	//
	//}
	//t1 := time.Now()
	//timestamp1 := t1.UnixNano() / 1000000 //UnitNano获取的是纳秒，除以1000000获取毫秒级的时间戳
	//
	//println("耗时：%d",timestamp1-timestamp)
	x := []int{100,110,120}
	for i, n:=range x  {
		println(i," n:",n)
	}
	y :=[]string{"I","have","a","lot","of","money"}
	for i,n:=range y {
		println(i,"n:",n)
	}
	//输出小于5的数字 相当于 while(i<5){}
	i :=0;
	for i<5 {
		println(i)
		i++;
	}


}

package main

import "fmt"
import "os"
import "time"

func main() {
	var s,sep string
	//var three int
	for i:=0;i<len(os.Args);i++{
		s+=sep+os.Args[i]
		sep=""
	//	three:=0;
	}
	fmt.Printf(s)
	for i:=0;i<100;i++{

	}
	//var num int
	//t1 := time.Now() // get current time
	//for{
	//	num++;
	//	fmt.Printf("%d \n",num)
	//	elapsed := time.Since(t1)
	//	if(elapsed>1000000000){
	//		fmt.Println("App elapsed: ", elapsed)
	//		break
	//	}
	//
	//}
	start := time.Now() // get current time
	sum:=0;
	for  i:=0;i<10000000;i++ {
		sum+=i;
	}
	consuming:=time.Since(start)
	fmt.Printf("\n耗时：%d \n",consuming.Nanoseconds())
	fmt.Printf("sum:%d",sum)

}

package main

import (
	"errors"
	"fmt"
)

func main() {
	//var e,d int
	//e,d=20,10
	////定义多个变量
	//a,b:=10,0
	////接收多个变量
	//c,err:=div(a,b)
	//println("println c err %v",c,err)
	//fmt.Println(c,err)
	//fmt.Println("E,D",e,d)
	//println("*****",e,d)
	//println("println")
	//fmt.Println("fmt.Println")
	var x,y int
	x,y=100,0
	//
	//f:=test(x)
	//f()
	de(x,y)
}

func div(a,b int)(int ,error)  {
	if b==0{
		return 0,errors.New("division by zero")
	}
	return a/b,nil
}

func test(x int) func()  { //返回函数类型
	return func() {//匿名函数
		fmt.Print(x) //闭包
	}
}
func de(a,b int)  {
	defer println(" dispose ...")    //常用来释放资源，释放锁定，或执行一些清理操作
	defer println(" dispose1 ...")   // 可以定义好多个defer,按filo顺序 first in last out 栈先进后出
	defer println(" dispose2 ...")
	println(a/b)

}

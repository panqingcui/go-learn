package main

import "fmt"
func main(){
  p := new(int)
  fmt.Printf("%d\n",p)
  fmt.Println("*p",*p)
  *p = 2
  fmt.Println("*p",*p)
  fmt.Println("newInt()",*newInt()) 
  v:=1
  fmt.Println("int1",newInt1())
 fmt.Println("v",v)
 fmt.Println("incr",*incr(&v))
}

func  newInt() *int{
   return new(int)
}
func   newInt1() int{
  var  dummy int
  return dummy
}

func incr(p *int) *int{
   *p++
   return p
}

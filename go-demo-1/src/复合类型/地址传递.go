package main

import "log"

func swap(p1, p2 *int) {
	*p1, *p2 = *p2, *p1
}
func main() {
	a, b := 10, 20
	swap(&a, &b)
	log.Printf("交换 地址传递---a = %v;b = %v", a, b)
}

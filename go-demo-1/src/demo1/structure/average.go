package main

import "fmt"

// 1 数组当作参数 使用 2 int / int 需要强转类型 3 变量都需要定义 或 通过 := 赋值 4/ 注意 形参是否指定 数组长度
func main() {
	fmt.Println("vim-go")
	var balance = [3]int{1, 2, 3}
	f := average(balance, 3)
	fmt.Println(f)
}
func average(arr [3]int, size int) float32 {
	var sum int
	for i := 0; i < size; i++ {
		sum += arr[i]
	}
	return float32(sum / size)

}

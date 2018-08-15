package main

import "fmt"

//二维数组  var array_name [x][y] variable_type
func main() {
	var a = [3][2]int{{0, 0}, {3, 2}, {2, 1}}
	fmt.Println("vim-go", a)
	var i, j int
	for i = 0; i < 3; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d]:%d \n", i, j, a[i][j])
		}
	}
}

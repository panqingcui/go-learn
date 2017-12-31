package main

import "fmt"

func main() {
	//创建 容量为5的切片
	x := make([]int, 0, 5)
	for i := 0; i < 8; i++ {
		//追加数据，当超出容量，自动分配更大的存储空间
		x = append(x, i)
	}
	fmt.Println(x)
	//创建字典型对象
	m := make(map[string]int)
	m["a"] = 1      //添加或设置
	y, ok := m["a"] //使用ok-idion 获取值，可知道key/value是否存在
	//所谓ok-idiom模式，是指在多返回值中用一个名为Ok的布尔值来表示操作是否成功，因为多操作默认返回零值，所以必须额外说明
	fmt.Println(y, ok)
	delete(m, "a")
	e, ok := m["a"] //使用ok-idion 获取值，可知道key/value是否存在
	fmt.Println(e, ok)

	var manage manage
	manage.name = "Tom"
	manage.age = 29
	manage.title = "CTO"
	fmt.Println(manage)

}

type user struct {
	name string
	age  byte
}
type manage struct {
	user
	title string
}

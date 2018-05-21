package main

import "fmt"

func main() {
	// map create :make(map[type of key]type of value) no init is :nil
	var persionSalary map[string]int
	if persionSalary == nil {
		fmt.Println("persionSalary is nil,going to make one ...")
		persionSalary = make(map[string]int)
		fmt.Println("persionSalary ", persionSalary)
		persionSalary["zhangsan"] = 1000000
		persionSalary["test"] = 2001020
		fmt.Println("persionSalary", persionSalary)
		test := map[string]int{
			"test": 123456789,
		}
		fmt.Println("test map :", test)
		fmt.Println("test map test key :", test["test"])
		// 存在判断
		newMap := "test"
		value, ok := persionSalary[newMap]
		if ok == true {
			fmt.Printf("Salary of %s value is %d", newMap, value)
		} else {
			fmt.Println("not found key ", newMap)
		}
		//循环变量

		for key, value := range persionSalary {
			fmt.Printf(" \npersionSalary key [%s] value is %d \n", key, value)
		}

	}
}

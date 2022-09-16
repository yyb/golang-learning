package main

import (
	"fmt"
)

func main() {
	type WeekDay int

	var Monday, Tuesday, Wednesday WeekDay

	Monday, Tuesday, Wednesday = 1, 2, WeekDay(3)

	fmt.Println(Monday, Tuesday, Wednesday)
	fmt.Printf("%T\n", Monday) //打印类型

	// 类型判断
	var value interface{}
	value = Tuesday

	// 是否WeekDay,这里的value是interface类型下的WeekDay，str得到的是纯WeekDay类型变量
	str, ok := value.(WeekDay)
	if ok {
		fmt.Printf("value is: %v\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
	// 通过switch中的特殊语法进行分类处理
	switch str := value.(type) {
	case WeekDay:
		fmt.Println("type WeekDay", str, value)
	case string:
		fmt.Println("type string")
	}

}

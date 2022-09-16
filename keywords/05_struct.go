package main

import "fmt"

type People struct {
	name string
	age  int
}

func main() {

	type Human struct {
		name string
		age  int
	}

	zhangsan := People{
		name: "张三",
		age:  20,
	}

	fmt.Println(zhangsan)

	var lisi struct {
		name string
		age  int
	}

	lisi.name = "李四"
	lisi.age = 20
	fmt.Println(lisi)

	// 结构一致，交叉赋值，不改变变量类型 list依然是临时结构
	lisi = zhangsan
	fmt.Println(lisi)

	lisi = Human{
		name: "王五",
		age:  100,
	}
	// 无法用 Human 结构 赋值给 People 结构的变量
	//zhangsan = Human{
	//	name: "王五",
	//	age:  100,
	//}

	fmt.Println(lisi)

	// 结构调用是复制
	changeName(zhangsan)
	fmt.Println(zhangsan)

	// 结构调用传引用
	changeNameRef(&zhangsan)
	fmt.Println(zhangsan)

	// 直接创建引用变量
	xiaohong := new(People)
	xiaohong.name = "小红"
	xiaohong.age = 15
	fmt.Println(xiaohong)
	changeNameRef(xiaohong)
	fmt.Println(xiaohong)
}

func changeName(zhangsan People) {
	zhangsan.name = "改名"
}

func changeNameRef(zhangsan *People) {
	zhangsan.name = "改名"
}

package main

import "fmt"

// map用于声明map类型数据
func main() {

	stringMap()

}

func stringMap() {
	// 初始化一个 map
	m := map[string]int{
		"one": 1,
		"two": 2,
	}
	m["three"] = 3
	m["four"] = 4

	fmt.Println(m)

	for key, value := range m {
		fmt.Println(key, value)
	}

	// 读取值
	choose := "one"
	value, ok := m[choose]
	fmt.Println(value, ok)

	choose = "five"
	value, ok = m[choose]
	fmt.Println(value, ok)

	// 删除值
	delete(m, "one")
}

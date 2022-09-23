package main

import (
	"encoding/json"
	"fmt"
)

/*
接口是一个抽象概念，表明可以实现的方向
结构与接口是没有关联的，接口变量用于探测
*/
func main() {
	// 声明一个接口变量，表示变量可以是任何实现
	var T interface{}

	T = 123
	fmt.Println(T)
	T = "abc"
	fmt.Println(T)
	T = struct {
		foo string
	}{"foo"}
	fmt.Println(T)

	var Y any = 123 //等价interface{}
	fmt.Println(Y)

	// json字段定义和输出
	type TestObj struct {
		Foo string `json:"tag 1"`
		Bar string `json:"tag 2"`
	}
	obj := TestObj{"字符类型", "没有"}
	obj.Foo = "替换词"
	res, _ := json.Marshal(obj)
	fmt.Println(string(res))

	book := new(HarryPotter)
	fmt.Println(book.isFun())
	fmt.Println(book.body())

	cat := new(Cat)
	// 进行接口定义
	var catInterface Emotion
	catInterface = cat
	_, hasEmotionInterface := catInterface.(Emotion)
	if hasEmotionInterface {
		fmt.Println("has emotion interface")
	}
	_, hasAnimalInterface := catInterface.(Animal)
	if hasAnimalInterface {
		fmt.Println("has animal interface")
	}
	fmt.Println(showEmotion(book))
	fmt.Println(showEmotion(catInterface))
	//catInterface.eat() 由于使用的是接口，所以无法调用接口定义外的方法
}

type (
	Emotion interface {
		isFun() bool
	}
	Story interface {
		Emotion
		body() string
		name() string
	}
	Animal interface {
		Emotion
		eat() string
	}
)
type HarryPotter struct {
	Book
}
type Cat struct {
}

func (c Cat) isFun() bool {
	return true
}

func (c Cat) eat() string {
	return "yummy"
}

type Book struct {
	bookName string
	bookBody string
}

func (p *Book) isFun() bool {
	return false
}

func (p *Book) body() string {
	return p.bookBody
}

func (p *Book) name() string {
	return p.bookName
}

func showEmotion(emotion Emotion) bool {
	return emotion.isFun()
}

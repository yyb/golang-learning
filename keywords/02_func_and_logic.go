package main

import "fmt"

/*
return ：用于从函数返回
func： 用于定义函数和方法
break、case、continue、for、fallthrough、else、if、switch、goto、default 流程控制
defer someCode ：在函数退出之前执行
*/
func main() {

	// func and return
	var result int
	result = Function()
	result1, result2 := Function1()
	fmt.Println(result, result1, result2)

	Function2(1, 2)
	Function2(3, 4, 5, 6)

	// 可变长参数
	args := []int{7, 8, 9}
	Function2(args...)

	// for
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	// for range
	for index, value := range []int{1, 2, 3} {
		fmt.Println(index, value)
	}

	//for break,label break, continue
Out:
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 {
				break //当前域break
			}
			if i == 2 && j > 0 {
				break Out //根据label break
			} else {
				fmt.Println(i, j)
				continue
			}

		}
	}

	//switch
	switchTest(1)
	switchTest(2)
	switchTest(3)

	//goto
	var loopTimes int
Loop:
	loopTimes++
	if loopTimes < 10 {
		goto Loop
	}

}

func switchTest(choose int) {
	defer TestDeferExit()    //最后执行
	defer TestDeferExitTwo() //优先于上一个defer执行
	switch choose {
	case 1:
		fmt.Println(1)
		break //有break语法
	case 2:
		fmt.Println(2)
		//但是没有break也不会继续执行下一个case
	case 3:
		fmt.Println(3)
		fallthrough // 强制执行下一个case
	case 4:
		fmt.Println(4)
	case 5, 6, 7, 8:
		fmt.Println(5, 6, 7, 8)
	default:
		fmt.Println("default")
	}
}

func TestDeferExitTwo() {
	fmt.Println("Test exit two")
}

func TestDeferExit() {
	fmt.Println("Test exit")
}

// 单值返回
func Function() int {
	return 123
}

// 多值返回
func Function1() (int, res int) {
	res = 2
	return 1, res
}

// 可变传入参数
func Function2(args ...int) {
	for index, value := range args {
		println("Function2", index, value)
	}
}

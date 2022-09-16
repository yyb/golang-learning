package main

import (
	"fmt"
)

/*
package and import: 导入
var和const ：变量和常量的声明
var varName type  或者 varName : = value
*/
func main() {
	// var 关键字声明变量
	var val int
	var (
		val1 int
		val2 int
	)
	var val3, val4 int
	var val5 func() bool
	val5 = func() bool {
		return false
	}
	var val6 struct {
		x int
	}
	var val7 interface{}

	type ValType int64

	var val8 ValType

	fmt.Println(val, val1, val2, val3, val4, val5(), val6, val7, val8)

	// const 关键字声明常量
	const (
		a = 1
		b
	)
	const c = 2
	const d byte = 97

	type ConstType int64

	const e ConstType = 123

	fmt.Println(a, b, c, d, e)

	//iota 常量生成器，生成顺序值枚举
	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	const (
		_          = iota // ignore first value by assigning to blank identifier
		KB float64 = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)

	fmt.Println(KB, MB, GB, TB, PB, EB, ZB, YB)

	// 自动推断变量类型
	number := 1
	fmt.Println(number)

	// 表达式推断可能会错误
	//maxUint := math.MaxUint
	// output:cannot use math.MaxUint (untyped int constant 18446744073709551615) as int value in assignment (overflows)
}

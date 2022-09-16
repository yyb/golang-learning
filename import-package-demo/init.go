package import_package_demo

import "fmt"

// 大写开头公共变量
var Gl int = 100

// 小写开头，私有变量
var private_gl = 200

func init() {
	fmt.Println("Gl:", Gl)
	fmt.Println("private_gl:", private_gl)
}

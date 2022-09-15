package main

import (
	"fmt"
	. "fmt"
	alias "fmt"
	"github.com/yyb/golang-learning/import-package-demo"   //导入并执行init方法
	_ "github.com/yyb/golang-learning/import-package-demo" //只执行init方法
)

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}
func init() {
	fmt.Println("init 3")
}

func main() {
	fmt.Println("引用包并使用")
	alias.Println("以别名引用包并使用")
	Println("导入包并使用")
	Println(import_package_demo.Gl)
}

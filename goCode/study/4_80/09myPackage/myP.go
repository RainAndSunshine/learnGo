package myP

import "fmt"

// import时 自动执行，init()初始化函数
func init() {
	fmt.Println("import我时自动调用myP的init()")
}

// 想被别的包调用的标识符都要首字母大写

// Add 加法函数
func Add(a, b int) int {
	return a + b
}

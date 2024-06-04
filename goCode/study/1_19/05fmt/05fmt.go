package main

import "fmt"

// func占位符
func main() {
	var n = 100
	//查看变量类型
	fmt.Printf("%T\n", n)
	//查看变量的值
	fmt.Printf("%v\n", n)
	//二进制
	fmt.Printf("%b\n", n)
	//八进制
	fmt.Printf("%o\n", n)
	//十进制
	fmt.Printf("%d\n", n)
	//十六进制
	fmt.Printf("%x\n", n)

	//字符串类型
	var s = "HelloGo"
	fmt.Printf("%s\n", s)
	//%v是万能取值
	fmt.Printf("%v\n", s)
	//加了井号，可以把具体的类型加上描述符，字符串会加上双引号
	fmt.Printf("%#v\n", s)
}

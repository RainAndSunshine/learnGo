package main

import "fmt"

//defer
//适用于释放资源，例如关闭文件、释放数据库的连接...
//防止忘记关闭文件，可以在打开文件后加上defer关闭文件

//一个函数中可以有多个defer语句
//与栈的先进后出一致，后defer的先执行

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("嘿嘿嘿") //defer把它后面的语句延迟到函数即将返回的时候再执行
	defer fmt.Println("哈哈哈") //与栈数据结构一致，先进后出
	fmt.Println("end")
}

func main() {
	deferDemo()
}

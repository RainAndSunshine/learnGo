package main

import "fmt"

// 变量的作用域
var x = 100 //定义一个全局变量
var y = 999

// 定义一个函数
func ff() {
	y := 10
	//函数中查找变量的顺序
	//1.先在函数内部查找
	//2.找不到就往函数的外面查找，直到找到全局
	fmt.Println(y) //函数内有，就用函数内的
	fmt.Println(x)
}

func main() {
	ff()
}

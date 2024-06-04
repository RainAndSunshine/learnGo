package main

import "fmt"

//// 匿名函数
//var ff1 = func(x, y int) {
//	fmt.Println(x + y)
//} //匿名函数一般不怎么用，一般都声明在函数内部

func main() {
	//函数内部无法声明一个有名字的函数
	//这时候就需要定义匿名函数
	ff1 := func(x, y int) {
		fmt.Println(x + y)
	}
	ff1(10, 20)

	//如果只是调用一次的函数，还可以简写成立即执行的函数
	func() {
		fmt.Println("Hello World!")
	}()
	//有参数的话
	func(x, y int) {
		fmt.Println(x, y)
	}(100, 200)

}

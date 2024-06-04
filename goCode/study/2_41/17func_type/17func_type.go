package main

import "fmt"

// 函数类型
// 函数类型还有区分
// 参数和返回值是类型的一部分
func f11() {
	fmt.Println("Hello World")
}

func f12() int {
	return 10
}

func f14(x, y int) int {
	return x + y
}

// 高阶函数，函数可以作为参数也可以作为返回值
// 函数也可以作为参数的类型
// 返回值为int，无参数的函数
func f13(x func() int) {
	ret := x()
	fmt.Println(ret)
}

// 函数也可以作为返回值
func f15(x func() int) func(int, int) int {
	ret := func(a, b int) int {
		return a + b
	}
	return ret
}

func main() {
	a := f11
	fmt.Printf("%T\n", a)
	b := f12
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", f14)
	f13(f12)
	f13(b)
	f15(f12)
}

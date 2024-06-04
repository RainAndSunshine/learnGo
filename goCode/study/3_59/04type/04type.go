package main

import "fmt"

//自定义类型和类型别名

//type 后面跟的是类型

// myInt 自定义类型
type myInt int

// yourInt 类型别名	ps：rune是int32的类型别名   byte是int8的类型别名
// 类型别名可以更方面分辨出应该输入什么数据，例如 rune
type yourInt = int

func main() {
	var n myInt
	n = 100
	fmt.Printf("type:%T Value:%v\n", n, n)
	var m yourInt = 100
	fmt.Printf("type:%T Value:%v\n", m, m)
	var c rune
	c = '陕'
	fmt.Printf("type:%T Value:%c\n", c, c)
}

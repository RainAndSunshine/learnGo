package main

import "fmt"

// 将var变成const则定义了常量
// 常量在定义之后不能修改
// 在程序运行期间不会改变
const pi = 3.1415926

// 多个常量也可以一起声明
const (
	n1 = 1
	n2 = 2
	n3 = 3
)

// const同时声明多个常量的时候，如果某一行声明后没有赋值，默认表示与上面一行相同
const (
	n4 = 3
	n5
	n6
)

// iota是go语言的常量计数器，只能在常量表达式中使用
// iota在const关键字出现时将被重置成0，const中每新增 一行 ，常量声明将使iota计数一次
const (
	a1 = iota
	a2 = iota
	a3 = iota
	a4 = iota
)

const (
	b1 = iota
	b2 = iota
	_  = iota
	b3 = iota
)

// 插队
const (
	c1 = iota
	c2 = 100
	c3 = iota
)

// 多个常量声明在一行
// 注意是每新增一行，iota+1
const (
	d1, d2 = iota + 1, iota + 2
	d3, d4 = iota + 1, iota + 2
)

// iota可以用于定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
)

func main() {
	fmt.Println(pi)
	fmt.Println("n1：", n1, "n2：", n2, "n3：", n3)
	fmt.Println("n4：", n4, "n5：", n5, "n6：", n6)
	fmt.Println(a1, a2, a3, a4)
	fmt.Println(b1, b2, b3)
	fmt.Println(c1, c2, c3)
	fmt.Println(KB, MB, GB, TB)
}

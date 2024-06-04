package main

import (
	"fmt"
	"math"
)

func main() {
	// 整型
	var i1 = 101
	fmt.Printf("%d\n", i1) // 十进制
	fmt.Printf("%b\n", i1) // 十进制转二进制
	fmt.Printf("%o\n", i1) // 十进制转八进制
	fmt.Printf("%x\n", i1) // 十进制转十六进制
	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2) // 十进制
	// 十六进制
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)
	// 查看变量的类型
	fmt.Printf("%T\n", i3)

	// 声明一个int8类型变量
	// 指定声明类型，否则会默认为int类型
	i4 := int8(9)
	fmt.Printf("%d\n", i4)
	fmt.Printf("%T\n", i4)

	// 浮点数
	mm := math.MaxFloat32
	fmt.Printf("%f\n", mm)
	f1 := 1.23456
	fmt.Printf("%T\n", f1) //默认go语言中都是float64
	//f1 = f2       //float32类型的值不能直接赋值给float64类型

	//复数，有实部和虚部
	//complex64的实部和虚部为32位的，complex128的实部和虚部为64位

	//布尔值
	isOk := true
	var is2 bool
	fmt.Printf("%T\n", isOk)
	//bool 默认为false
	fmt.Printf("%T value:%v\n", is2, is2) //%v打印变量的值

}

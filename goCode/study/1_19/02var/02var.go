package main

import "fmt"

// 全局变量不使用也可以编译通过
// 非全局变量声明后必须使用
// 声明变量
var s1 string
var age int
var isOk bool

// 批量声明
var ( //默认空值
	name string //""
	age1 int    //0
	//Go语言中推荐使用小驼峰式命名
	isOk1 bool //false
)

func main() {
	name = "理想"
	age1 = 16
	isOk1 = true
	// Go语言中的变量声明必须使用，不使用就编译不过去，避免出现多余变量
	// 全局变量不使用也可以编译通过
	// 非全局变量声明后必须使用
	fmt.Printf("name:%s\n", name)
	fmt.Println(age)
	fmt.Printf("isOk:%v", isOk1)
	fmt.Println()
	// 声明变量同时赋值
	var s11 string = "whb"
	fmt.Println(s11)
	// 类型推导
	var s2 = "逆天"
	fmt.Println(s2)
	// 简短变量声明，只能在函数里面使用，不能在全局使用
	s3 := "哈哈哈"
	fmt.Println(s3)
	// 匿名变量，用一个单独的下划线“_”表示
	// 在使用多重赋值的时候，如果需要忽略某一个值，则使用匿名变量
	x, _ := 15, "牛逼"
	_, y := 15, "牛逼"
	fmt.Println(x, y)

}

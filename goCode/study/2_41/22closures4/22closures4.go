package main

import "fmt"

//闭包 = 函数 + 外部变量的引用

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		//base是这些函数外部的，引用参数
		//修改的是外部的base
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	f1, f2 := calc(10)
	//base是一个引用，修改后，base也会跟着变化
	//修改的是外面公用的base
	//执行了base就会随着变化
	fmt.Println(f1(1), f2(2)) //11 9
	fmt.Println(f1(3), f2(4)) //12 8
	fmt.Println(f1(5), f2(6)) //13 7
}

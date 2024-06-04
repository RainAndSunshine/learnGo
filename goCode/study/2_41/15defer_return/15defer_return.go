package main

import "fmt"

//defer执行时机
//在Go语言的函数中return语句在底层并不是原子操作，它分为 给返回值赋值 和 RET指令
//defer语句的执行时机在 给返回值赋值之后  RET指令之前
//即
//1.给返回值赋值
//2.defer语句
//3.RET指令

// 函数用的是未命名的返回值
// 先给的返回值赋值为5
// 后运行的x++，这个x与返回值无关
// 输出的是5
func a1() int {
	x := 5
	defer func() {
		x++ //修改的是x，不是返回值
	}()
	return x
}

// 函数用的是命名的返回值
// 先给返回值赋值5
// 随后执行x++，这个x就是返回值，即返回值++
// 输出的是6
func a2() (x int) {
	defer func() {
		x++ //返回值x++
	}()
	return 5
}

// 函数用的是命名的返回值
// 但命的名与return的不同
// 先给返回值y赋值，相当于执行了一句y=x
// 后执行x++，与返回值y无关
// 输出的是5
func a3() (y int) {
	x := 5
	defer func() {
		x++ //修改的是x，不是返回值
	}()
	return x //相当于执行了一句y=x
}

// 函数用的是命名的返回值
// 但里面的修改语句是一个函数
// 将x作为参数，函数传参是一个拷贝，不影响原值
// 输出的是5
func a4() (x int) {
	defer func(x int) {
		x++ //改变的是x的副本
	}(x) //函数传参是拷贝，不影响原值
	return 5 //返回值 =x =5
}

func a5() (x int) {
	defer func(x int) int {
		x++      //改变的是x的副本
		return x //return的结果没人接收,没有效果
	}(x) //函数传参是拷贝，不影响原值
	return 5 //返回值 =x =5
}

// 传一个x的指针到匿名函数中
func a6() (x int) {
	defer func(x *int) {
		(*x)++ //改变的是x地址指向的值，所以改变的就是原来的值
	}(&x) //传的是x的地址，地址不管怎么复制，都是一样的
	return 5 //1.返回值=x=5  2.defer x=6  3.RET返回 6
}

func main() {
	fmt.Println(a1()) //5
	fmt.Println(a2()) //6
	fmt.Println(a3()) //5
	fmt.Println(a4()) //5
	fmt.Println(a5()) //5
	fmt.Println(a6()) //6
}

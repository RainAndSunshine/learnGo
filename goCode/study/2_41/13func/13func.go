package main

import "fmt"

//函数
//是一段代码的封装

// 函数的定义
// func + 函数名 + (参数) + (返回值)
func my_Sum(x int, y int) (sum int) {
	return x + y
}

// 没有返回值
func my_sum1(x1 int, x2 int) {
	fmt.Println(x1 + x2)
}

// 没有参数没有返回值
func my_sum2() {
	fmt.Println(3)
}

// 没有参数但有返回值
func my_sum3() int {
	return 3
}

// 返回值可以命名也可以不命名
// 要么都命名，要么都不命名
func my_sum(x int, y int) (ret int) {
	//命名了就可以直接使用，命名的返回值就相当于在函数中声明了一个变量
	//命名了的返回值，return的时候可以不写名称，默认return该变量
	ret = x + y
	return
}

// 多个返回值
func f1() (int, string) {
	return 1, "可以"
}

// 参数的类型简写，当参数中连续多个参数的类型一致时，可以将前面的参数类型省略
func f2(x, y, z, j, k int, n, m string, n1, m1 bool) int {
	return x + y
}

// 可变长参数，是一个切片
// 可变长参数必须放在函数的最后
func f3(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) //y的类型是切片，[]int
}

//Go语言中函数没有默认参数的概念

func main() {
	//函数的调用
	r := (my_Sum(10, 20))
	fmt.Println(r)
	a1, a2 := f1()
	fmt.Println(a1, a2)
	_, b := f1()
	fmt.Println(b)
	a, _ := f1()
	fmt.Println(a)

	//用省略号的，可以传多个
	f3("北京", 1, 2, 3, 4, 5)
	f3("北京", 1, 2, 3)
	//也可以不传
	f3("北京")
}

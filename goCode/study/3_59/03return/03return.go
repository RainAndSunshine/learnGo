package main

import "fmt"

// 递归：函数自己调用自己
//递归适合处理
// 递归必须要有一个明确的退出条件

// 使用递归计算阶层
func f1(a int) int {
	if a <= 1 {
		return 1
	}
	return a * f1(a-1)
}

// 跳台阶
func taijie(a int) int {
	if a == 1 {
		return 1
	}
	if a == 2 {
		return 2
	}
	return taijie(a-2) + taijie(a-1)
}

func main() {
	var res int
	fmt.Scan(&res)
	fmt.Println(f1(res))
	fmt.Scan(&res)
	fmt.Println(taijie(res))
}

package main

import "fmt"

// 要求
// t1(t2)

func t1(t func()) {
	fmt.Println("this is t1")
	t()
}

func t2(x, y int) {
	fmt.Println("this is t2")
	fmt.Println(x + y)
}

// 闭包
func t3(tt func(int, int), x, y int) func() {
	fmt.Println("this is t3")
	ret := func() {
		tt(x, y)
	}
	return ret
}

func main() {
	//有了闭包，t2就可以作为参数传入t1中了
	tep := t3(t2, 100, 200)
	t1(tep)
}

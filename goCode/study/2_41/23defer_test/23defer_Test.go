package main

import "fmt"

//defer

func calc1(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	//要先将函数内的参数计算出来，所以到这里就执行calc1("A", x, y)
	//而且函数内的变量都会被赋值，不能留有变量，应该都为确定的量
	//所以这个defer 的x  =   1
	defer calc1("AA", x, calc1("A", x, y))
	x = 10
	//这里也一样，先执行calc1("B", x, y)
	defer calc1("BB", x, calc1("B", x, y))
	y = 20
	//结束之后，按先进后出的顺序
	//先执行calc1("BB", x, calc1("B", x, y))
	//后执行calc1("AA", x, calc1("A", x, y))
}

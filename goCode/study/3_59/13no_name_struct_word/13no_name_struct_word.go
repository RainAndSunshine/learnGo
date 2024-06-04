package main

import "fmt"

//匿名字段
//字段比较少，也比较简单的场景，且不常用

type person3 struct {
	string
	int
}

func main() {
	p1 := person3{
		"周林",
		9000,
	}
	fmt.Println(p1)
	//如果要取其中的值，匿名字段 默认为它的数据类型
	//所以相同数据类型只能写一个，不然不唯一，有冲突
	fmt.Println(p1.string)
	fmt.Println(p1.int)
}

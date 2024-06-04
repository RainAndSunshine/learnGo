package main

import "fmt"

type person2 struct {
	name string
	age  int
}

// 构造函数：约定成俗用new开头
// 当结构体比较大的时候，尽量使用结构体指针，减少程序的内存开销
func newPerson(name string, age int) *person2 { //因为函数是值拷贝且结构体也是值类型，赋值的时候都是拷贝
	//如果结构体内部字段很多的时候，多个拷贝会造成效率慢，浪费内存空间
	return &person2{
		name: name,
		age:  age,
	}
}

func main() {
	p01 := newPerson("John", 18)
	p02 := newPerson("周林", 9000)
	fmt.Println(p01, p02)
}

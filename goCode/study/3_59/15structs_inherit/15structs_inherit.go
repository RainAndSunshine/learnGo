package main

import "fmt"

//Go语言中没有继承这个概率
//但可以模拟继承
//Go语言中使用结构体也可以模拟实现其他编程语言中面向对象的继承

type animal struct {
	name string
}

// 给animal实现一个移动的方法
func (a animal) move() {
	fmt.Printf("%s谁会动!\n", a.name)
}

// 狗类
type dog1 struct {
	feet   uint8
	animal //animal拥有的方法，dog此时也有了
}

// 给dog实现一个汪汪汪的方法
func (d dog1) barking() {
	fmt.Printf("%s在叫：汪汪汪~\n", d.name)
}

func main() {
	d1 := dog1{
		feet: 4,
		animal: animal{
			name: "zhoulin",
		},
	}
	fmt.Println(d1)
	d1.barking()
	d1.move() //move之所以能够执行，因为dog这个结构体包含了animal
}

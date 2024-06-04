package main

import "fmt"

//使用值接受者和指针接受者的区别？
//使用值接受者实现接口，结构体类型和结构体指针类型的变量都能存
//指针接受者实现接口只能存结构体指针类型的变量

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

//// 使用值接受者实现了接口的所有方法
//func (c cat) move() {
//	fmt.Println("走猫步")
//}
//
//func (c cat) eat(name string) {
//	fmt.Printf("猫吃%s\n", name)
//}

// 使用指针接受者实现了接口的所有方法
func (c *cat) move() {
	fmt.Println("走猫步")
}

func (c *cat) eat(name string) {
	fmt.Printf("猫吃%s\n", name)
}

func main() {
	var a1 animal
	c1 := cat{"Tom", 4}  //cat
	c2 := &cat{"假老练", 4} //*cat

	a1 = &c1 //实现animal这个接口的是cat的指针类型，需要取地址符号来取指针
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)
}

package main

import "fmt"

//接口的定义
//type 接口名 interface {
//	方法名1(参数1, 参数2...)(返回值1,返回值2...)
//	方法名2(参数1, 参数2...)(返回值1,返回值2...)
//	...
//}
//用来给变量\参数\返回值等设置类型

//接口的实现
//一个变量如果实现了接口中规定的所有的方法，那么这个变量就实现了这个接口，可以称为这个接口类型的变量
//即可以将实现接口所有方法的类型，赋值给接口类型

type animal interface {
	move()
	eat(string) //在接口里可以不写参数名，只要写参数的类型就好了
}

type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("走猫步")
}

func (c cat) eat(food string) {
	fmt.Printf("猫吃%s\n", food)
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("鸡动")
}

func (c chicken) eat(food string) {
	fmt.Printf("鸡吃%s\n", food)
}

func main() {
	var a1 animal //定义一个接口类型的变量
	//没赋值的接口，类型是nil
	//因为接口的储存分为两部分
	//1.动态类型
	//2.动态值
	//如果没赋值，接口储存的动态类型和动态值都是nil
	//接口是一种引用类型，抽象的，没有实际的东西，只是一个约束，用来约束变量
	fmt.Printf("%T\n", a1)

	bc := cat{ //定义一个cat类型的变量
		name: "蓝猫",
		feet: 4,
	}

	kfc := chicken{
		feet: 2,
	}

	var a2 animal = kfc
	fmt.Println(a2)
	a2.move()
	fmt.Printf("%T\n", a2)
	//接口的储存分为两部分
	//1.值的类型（动态类型）
	//2.值本身（动态值）
	//这样就实现了接口变量能够存储不同的值

	a1 = bc
	a1.eat("小黄鱼")
	fmt.Println(a1)
	fmt.Printf("%T\n", a1)
}

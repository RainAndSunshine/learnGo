package main

import "fmt"

//接口和类型的关系
//多个类型可以实现同一个类型
//一个类型可以实现多个接口

// 同一个结构体可以实现多个接口
//接口还可以嵌套

// 接口的嵌套
type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// cat实现了mover接口
func (c *cat) move() {
	fmt.Println("走猫步")
}

// cat还实现了eater接口
func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s\n", food)
}

func main() {

}

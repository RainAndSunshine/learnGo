package main

import "fmt"

//接口示例
//不管是什么品牌的车，都能跑

// 定义一个car接口类型
// 不管是什么结构体，只要有run方法都能car类型
type car interface {
	run()
}

type ferrari struct {
	brand string
}

type porsche struct {
	brand string
}

func (f ferrari) run() {
	fmt.Printf("%s速度99999迈\n", f.brand)
}

func (p porsche) run() {
	fmt.Printf("%s速度999迈\n", p.brand)
}

var f1 = ferrari{
	brand: "ferrari",
}

var p1 = porsche{
	brand: "porsche",
}

func drive(c car) {
	c.run()
}

func main() {
	drive(f1)
	drive(p1)
}

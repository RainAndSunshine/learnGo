package main

import "fmt"

//结构体，可以将一个事物的全部信息都保存起来

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var p1 person
	//通过字段赋值
	p1.name = "lixiang"
	p1.age = 18
	p1.gender = "男"
	p1.hobby = []string{"篮球", "双色球"}
	fmt.Println(p1)
	//访问变量p1的字段
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.name)

	//匿名结构体：多用与临时场景
	var s struct {
		x string
		y int
	}
	s.x = "嘿嘿嘿"
	s.y = 100
	fmt.Printf("type:%T value:%v\n", s, s)

}

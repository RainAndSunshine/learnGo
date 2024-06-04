package main

import "fmt"

//结构体是值类型

type person1 struct {
	name, gender string
}

// go语言中函数传的参数永远是拷贝
func f(x person1) {
	x.gender = "女" //修改的是副本的gender
}

func ff(x *person1) {
	// (*x).gender = "女"
	x.gender = "女" //可以直接写成这个，因为go语言是不可以对指针进行修改的，会自动根据指针找对应的量
	//根据内存地址找到那个原变量，修改的就是那个原变量
}

func main() {
	//值类型的结构体
	var p person1
	p.name = "李文周"
	p.gender = "男"
	f(p)
	fmt.Println(p.gender) //还是男
	ff(&p)
	fmt.Println(p.gender) //变成女了

	//结构体指针
	//创建一个指针类型的结构体
	var a = new(person1)
	a.name = "理想"
	fmt.Printf("%T\n", a)  //指针类型
	fmt.Printf("%p\n", a)  //a的值就是一个内存地址
	fmt.Printf("%p\n", &a) //求a的内存地址

	//结构体初始化
	//初始化的时候也可以直接创建成一个结构体指针
	//1） key-value初始化
	var b = &person1{
		name:   "元帅",
		gender: "男", //这里的逗号不能少
	}
	fmt.Printf("%#v\n", b)

	//2) value直接初始化
	//值的顺序必须和结构体定义时字段的顺序一致
	var c = &person1{
		"原神",
		"男",
	}
	fmt.Printf("%#v\n", c)
}

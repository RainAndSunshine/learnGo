package main

import "fmt"

//结构体嵌套

type address struct {
	province string
	city     string
}

type workPlace struct {
	province string
	city     string
}

type person4 struct {
	name string
	age  int
	addr address
	workPlace
}

type company struct {
	name      string
	address   //匿名嵌套结构体
	workPlace //当出现两个匿名嵌套结构体，访问的时候可能会出现冲突
}

func main() {
	p1 := person4{
		name: "秀才",
		age:  9000,
		addr: address{
			province: "山东",
			city:     "济南",
		}, //这个逗号不能忘
		workPlace: workPlace{ //匿名嵌套结构体的初始化
			province: "湖北",
			city:     "武汉",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.name)
	//嵌套结构体的访问
	fmt.Println(p1.addr.city)
	//没有直接叫city的，所以转去匿名嵌套结构体里找
	fmt.Println(p1.city) //先在自己结构体找这个字段，找不到就去匿名嵌套的结构体中查找该字段

	//如果需要跟之前一样可以直接访问，可以使用匿名嵌套结构体，匿名字段
	c1 := company{
		name: "腾讯",
		address: address{
			province: "广东",
			city:     "深圳",
		},
	}
	fmt.Println(c1)
	//使用了匿名嵌套结构体后，两种访问方式都可以
	fmt.Println(c1.address.city)
	//fmt.Println(c1.city) //先在自己结构体找这个字段，找不到就去匿名嵌套的结构体中查找该字段
	//因为两个匿名嵌套结构体都有city，所以出现冲突，只能具体写
	fmt.Println(c1.address.city)
	fmt.Println(c1.workPlace.city)
}

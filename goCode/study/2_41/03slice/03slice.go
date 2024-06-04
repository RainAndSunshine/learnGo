package main

import "fmt"

func main() {
	//切片slice
	//数组的长度是固定不变的并且数组长度属于类型的一部分,数组是一个值类型
	//切片是一个拥有相同类型元素的可变长度的序列，它是基于数组类型做的一层封装
	//切片是一个引用类型，它的内部结构包含：地址  长度  容量

	//切片的定义
	var s1 []int //定义一个存放int类型的切片
	var s2 []string
	fmt.Println(s1, s2)
	//nil就是null，是空
	fmt.Println(s1 == nil, s2 == nil)

	//初始化1
	s1 = []int{1, 2, 3}
	s2 = []string{"Hello", "Kitty", "World"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil, s2 == nil)
	//初始化2
	//由数组得到切片
	a1 := [...]int{1, 3, 4, 5, 6, 7, 14}
	s3 := a1[0:4] //基于一个数组切割    左包含,右不包含（左闭右开）
	fmt.Println(s3)
	s4 := a1[1:6]
	fmt.Println(s4)
	s5 := a1[:4] //从头切到4    不包含4[0:4]
	fmt.Println(s5)
	s6 := a1[3:] //从4切到尾	  包含尾[3:len(a1)]
	fmt.Println(s6)
	s7 := a1[:] //从头切到尾	  全都要[0:len(a1)]
	fmt.Println(s7)

	//长度和容量
	//长度<=容量
	//切片的容量是底层数组从切片的第一个元素到最后一个元素数量
	//切片是支持扩容的
	fmt.Printf("len(s1):%d , cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d , cap(s2):%d\n", len(s2), cap(s2))

	//切片再切片
	//切片的容量是底层数组从切片的第一个元素到最后一个元素数量
	//这是一个数组类型，切片的[]里面是空的
	a2 := [...]string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
	fmt.Printf("a:%v type:%T len:%d cap:%d\n", a2, a2, len(a2), cap(a2))
	b2 := a2[1:3]
	fmt.Printf("b2:%v type:%T len:%d cap:%d\n", b2, b2, len(b2), cap(b2))
	c2 := b2[1:5]
	fmt.Printf("c2:%v type:%T len:%d cap:%d\n", c2, c2, len(c2), cap(c2))

	//切片是引用类型，都指向了底层的一个数组
	d2 := a2[2:]
	//切片的切片指向同一个底层数组
	e2 := d2[3:]
	fmt.Println("改之前e2：", e2)
	fmt.Println("改之前d2：", d2)
	a2[5] = "广东省广州市"
	fmt.Println("改之后d2：", d2)
	fmt.Println("改之后e2：", e2)
}

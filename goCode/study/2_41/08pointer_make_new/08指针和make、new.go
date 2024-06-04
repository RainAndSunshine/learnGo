package main

import "fmt"

func main() {
	//指针
	//Go语言中不存在指针操作，只需要记住两个符号
	//'&':取地址
	//'*':根据地址取值
	n := 18
	p := &n
	fmt.Printf("%T\n", p) //*int：int类型的指针
	fmt.Println(p)
	m := *p
	fmt.Printf("%T\n", m) //int类型
	fmt.Println(m)

	//对变量进行取地址（&）操作，可以获得这个变量的指针变量
	//指针变量的值是指针地址
	//对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值

	//使用new去创建指针变量
	//new只适用于基本数据类型
	var a = new(int)
	*a = 10
	fmt.Println(*a)

	//make区别于new，它只适用于slice、map以及channel的内容创建
	//返回的类型就是这三个类型本身，因为它们都是引用类型，没必要返回指针
	s1 := make([]int, 5, 10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	//make和new的区别
	//1.make和new都是用来申请内存的
	//2.new很少用，一般用来给基本数据类型申请内存,string\int...,返回的是对应类型的指针(*string\*int...)
	//3.make是用来给slice、map、channel申请内存的，make函数返回的是对应的这三个类型本身
}

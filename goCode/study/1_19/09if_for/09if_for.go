package main

import (
	"fmt"
)

func main() {
	age := 19
	if age >= 18 {
		fmt.Println("已满18岁")
	} else {
		fmt.Println("未满18岁")
	}
	//多个条件判断
	if age > 18 {
		fmt.Println("已成年")
	} else if age == 18 {
		fmt.Println("刚成年")
	} else {
		fmt.Println("未成年")
	}

	//在条件判断前可以增加一条执行语句
	//不过执行语句的作用域只在if内
	if age1 := 19; age1 >= 18 {
		fmt.Println("成年咯，嘿嘿嘿")
	} else {
		fmt.Println("成年人的世界你不懂")
	}
	//age1在if的作用域内定义，出了作用域就无效了
	//fmt.Println(age1)

	//Go语言中只有for一种循环
	//基本语法
	//for 初始语句; 条件表达式; 结束语句 {
	//	循环体语句
	//}
	//基本语法
	for i := 0; i < 10; i++ {
		fmt.Print(i, "\t")
		if i == 9 {
			fmt.Println()
		}
	}
	//初始值可以省略，但是分号必须写
	t := 0
	for ; t < 10; t++ {
		fmt.Print(t, "\t")
		if t == 9 {
			fmt.Println()
		}
	}
	//结束语句也可以省略
	var i = 5
	for i < 10 {
		fmt.Print(i, "\t")
		i++
	}
	fmt.Println()
	//直接的无限循环，可以通过break，goto，return，panic语言强制退出循环
	count1 := 0
	for {
		fmt.Print("死循环")
		count1++
		if count1 > 10 {
			break
		}
	}
	fmt.Println()
	//键值循环for range
	//Go语言中可以使用for range遍历数组、切片、字符串、map及通道（channel）
	//1.数组、切片、字符串返回索引和值
	//2.map返回键和值
	//3.通道（channel）只返回通道内的值
	s := "Hello祖国"
	//字符串返回索引和值
	for i, v := range s {
		fmt.Printf("%d %c\t", i, v) //中文占三个字节的空间
	}
}

package main

import "fmt"

func main() {
	//数组array

	//存放元素的容器
	//必须指定存放的容器的类型和容量（长度）
	//数组的长度是数组类型的一部分（长度和容量只要有一个不一样，两个数组就不一样，不可以进行比较）
	var a1 [3]bool
	var a2 [4]bool
	fmt.Printf("a1:%T \t a2:%T \n", a1, a2) //a1:[3]bool       a2:[4]bool

	//数组的初始化
	//如果不初始化，数组的元素默认都是零值（布尔值：false  整型和浮点型：0  字符串：""）
	fmt.Println(a1, a2)

	//方式1
	a1 = [3]bool{true, false, true}
	fmt.Println(a1)

	//方式2
	//根据初始值自动推断数组的长度是多少
	a100 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a100)

	//方式3
	//根据索引进行初始化
	a5 := [5]int{0: 1, 4: 2}
	fmt.Println(a5)

	//数组的遍历
	cities := [...]string{"北京", "上海", "广州", "深圳"}
	//1.索引遍历
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}
	//2.for range遍历
	//索引和值
	for i, v := range cities {
		fmt.Println(i, v)
	}

	//多维数组
	//在一个数组里只能放相同类型的
	//[[1 2] [3 4] [5 6]]
	//数组里面有三个元素，元素是包含两个元素数组
	var a11 [3][2]int

	//初始化多维数组
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)
	//多维数组，外层的可以...内层不可以
	a22 := [...][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a22)

	//多维数组的遍历
	for _, v := range a11 { //遍历出来的每一项又是数组
		fmt.Println(v)
		for _, w := range v { //遍历第一次的结果，取出元素
			fmt.Println(w)
		}
	}

	//数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)
}

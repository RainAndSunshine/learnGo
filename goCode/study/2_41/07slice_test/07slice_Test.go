package main

import (
	"fmt"
	"sort"
)

func main() {
	//说出下列程序的运行结果
	//长度为5，容量为10
	a := make([]int, 5, 10)
	fmt.Println(a)
	//append是追加，原来的长度为5，已经有5个元素在前面了，默认为0
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)

	//利用内置的sort包对数组var a1=[...]int{3,7,8,9,1}进行排序
	var a1 = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a1[:])
	fmt.Println(a1)
}

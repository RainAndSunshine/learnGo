package main

import (
	"fmt"
)

func main() {
	//map
	//无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用
	//map[key的类型]值的类型
	//一定要记得初始化
	m1 := make(map[string]int, 10) //要估算好map的容量，避免在程序运行期间，再次扩容
	m1["Hello"] = 100
	m1["World"] = 15
	fmt.Println(m1)
	fmt.Println(m1["Hello"])

	//如果不存在这个key拿到对应类型的零值
	fmt.Println(m1["没有"])
	//约定成俗，用ok接收返回的bool值
	value, ok := m1["没有"]
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(value)
	}

	//map 的遍历
	for key, value := range m1 {
		fmt.Println(key, value)
	}
	fmt.Println("nbnbnb")
	for key, value := range m1 {
		fmt.Println(key, &value)
	}
	//如果只想遍历key，只用写一个就行
	for k := range m1 {
		fmt.Println(k)
	}
	//如果只想遍历value，使用匿名变量
	for _, v := range m1 {
		fmt.Println(v)
	}

	//删除delete
	//按照key值删除
	delete(m1, "Hello")
	fmt.Println(m1)
	delete(m1, "没有") //不会报错，但不会进行操作

}

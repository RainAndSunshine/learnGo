package main

import "fmt"

//空接口
//没必要用type来定义空接口，通常直接interface{}
//type xx interface{
//
//}

//因为空接口没有必要起名字，所以通常定义为下面这种格式：
//interface{}	//常用的空接口
//所有的类型都实现了空接口，也就是任意类型的变量都能保存到空接口中
//有些函数想实现不管什么类型都可以传进来，参数就应该设置成空接口类型
//例如Println，map[key]value...

// 空接口作为函数的参数
func show(a interface{}) {
	fmt.Printf("type: %T, value: %v\n", a, a)
}

func main() {
	//interface：	关键字
	//interface{}：	空接口类型

	//空接口作为map的值
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "周林"
	m1["age"] = 9000
	m1["married"] = true
	//这是一个数组类型，切片的[]里面是空的
	m1["hobby"] = [...]string{"唱", "跳", "rap"}
	fmt.Println(m1)

	show(false)
	show(nil)
	show(m1)
}

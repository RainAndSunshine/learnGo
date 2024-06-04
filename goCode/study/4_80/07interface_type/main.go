package main

import "fmt"

//类型断言
//用来判断空接口中的值
//格式：x.(T)		其中：x表示类型为interface{}的变量，T表示断言x可以是的类型

//我想知道空接口接受的值具体是什么，就用类型断言

// 类型断言1
func assign(a interface{}) {
	//这个只能展示接口的类型，并不能拿到
	fmt.Printf("%T\n", a)
	//使用类型断言
	str, ok := a.(string)
	if ok {
		fmt.Println("传进来的是一个字符串", str)
	} else {
		fmt.Println("猜错了")
	}
}

// 类型断言2
func assign2(a interface{}) {
	//这个只能展示接口的类型，并不能拿到
	fmt.Printf("%T\n", a)
	//使用类型断言
	switch t := a.(type) {
	case string:
		fmt.Println("是一个字符串:", t)
	case int:
		fmt.Println("是一个int:", t)
	case int64:
		fmt.Println("是一个int64:", t)
	case bool:
		fmt.Println("是一个布尔值:", t)
	default:
		fmt.Println("错误猜测")
	}
}

func main() {
	assign(100)

	assign2(true)
	assign2("哈哈哈")
	assign2(int64(200))
}

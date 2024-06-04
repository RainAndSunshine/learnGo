package main

import "fmt"

//内置函数

//close：主要用来关闭channel
//len：用来求长度，比如string、array、slice、map、channel
//new：用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针
//make：用来分配内存，主要用来分配引用类型，比如chan、map、slice
//append：用来追加元素到数组、slice中
//panic和recover：用来做错误处理

//panic 和 recover成对出现

func funcA() {
	fmt.Println("a")
}

func funcB() {
	//刚刚打开了数据库连接
	//出错了要释放数据库连接
	//这时候可以用到defer
	defer func() {
		//在panic之后，用recover可以尝试恢复一下
		err := recover() //recover尽量别用
		fmt.Println(err)
		fmt.Println("释放数据库连接...")
	}()
	panic("出现了严重的错误！！！") //程序崩溃退出
	fmt.Println("b")
}

func funcC() {
	fmt.Println("c")
}

func main() {
	funcA()
	funcB()
	funcC()
}

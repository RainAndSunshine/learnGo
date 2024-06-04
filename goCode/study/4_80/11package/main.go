package main

import (
	"fmt"

	//调用时省略包名，但不建议怎么用，会造成标志符冲突
	//而且会弄混各个标志符
	. "goCode/study/4_80/10myPackage2"

	//起别名
	sust "goCode/study/4_80/09myPackage"
	//调用时省略包名
)

//Go语言禁止循环导入包

//如果只希望导入包，而不使用内部的数据时，可以使用匿名导入包
//只使用包里的init()初始化函数
//格式：import _ "包的路径"

// init()初始化函数
// 没有参数也没有返回值
// 一个包里只能定义一个init()
// 不能手动调用，只能自动调用
func init() {
	fmt.Println("import我时自动调用main的init()")
	fmt.Println(a, b, c)
}

// 先进行全局声明
// 然后执行init()初始化函数
// 最后执行main()函数
var a = 100
var b = "全局变量先声明"

const c = 666

//导入包顺序			main -> import -> A -> import -> B -> import -> C
//init()执行顺序		C.init() -> B.init() -> A.init() -> main.init()

func main() {
	res := sust.Add(100, 200)
	fmt.Println(res)
	//因为上面用了import . "路径名"
	//这种方式调用，不能写包名
	Mysum(100, 200, 300, 400)
}

package main

import (
	"fmt"
)

func main() {
	//编写代码分别定义一个 整型，浮点型，布尔型，字符串型 变量，使用fmt.Printf()搭配%T分别打印出上述变量的值和类型
	var a int64
	a = 10
	var b float64
	b = 3.14
	c := true
	s := "Hello美丽西安"
	fmt.Printf("a:%T,a=%d\n", a, a)
	fmt.Printf("b:%T,b=%f\n", b, b)
	fmt.Printf("c:%T,c=%v\n", c, c)
	fmt.Printf("s:%T,s=%s\n", s, s)

	//统计字符串中汉字的数量
	count := 0
	for _, cc := range s {
		//超纲
		//if unicode.Is(unicode.Han, cc) {
		//	count++
		//}

		if cc > 'z' {
			count++
		}
	}
	fmt.Printf("count:%d\n", count)
}

package main

import "fmt"

func main() {
	//byte和rune类型
	//Go语言中为了处理非ASCII码类型的字符(中文或者什么字符)，定义了新的rune类型
	s := "Hello西安"
	n := len(s)
	fmt.Println(n)
	for _, c := range s {
		fmt.Printf("%c\n", c)
	}
	//字符串修改
	s2 := "白萝卜"
	fmt.Println(s2)
	//字符串是不能修改的变量
	//s2[0] = "红"
	//如果要修改，需要把字符串强制转换为rune类型
	s3 := []rune(s2)        //把字符串强制转换成了一个rune切片
	s3[0] = '红'             //只修改一个字符，所以是字符型不是字符串型
	fmt.Println(string(s3)) //把rune切片强制转换为字符串
	c1 := "红"
	c2 := '红'
	//rune其实是int32的别名
	fmt.Printf("c1:%T , c2:%T\n", c1, c2)
	c3 := "H"
	c4 := 'H'
	fmt.Printf("c3:%T , c4:%T\n", c3, c4)
	fmt.Printf("c4:%d\n", c4) //输出对应的ASCII值

	//类型转换
	n1 := 10      //int
	var f float64 //float64
	f = float64(n1)
	fmt.Println(f)
	fmt.Printf("f:%T\n", f)
}

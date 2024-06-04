package main

import (
	"fmt"
	"strings"
)

func main() {
	//Go语言字符串是用双引号包裹的
	//Go语言字符才是用单引号包裹的

	s1 := "hello陕科大"
	c1 := 'h'
	c2 := '1'
	c3 := '陕'
	fmt.Println(s1)
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	//多行的字符串，用反引号，Esc下面的`
	//按照反引号里面的格式原样输出
	s2 := `
	世情薄
		人情恶
雨送黄昏花易落
`
	fmt.Print(s2)

	//字符串的常用操作
	s3 := "123.456.789"

	//计算字符串的长度
	fmt.Println(len(s3))
	s4 := "理想"

	//字符串的拼接
	//直接用加号+进行拼接
	ss := s3 + s4
	fmt.Println(s3 + s4)
	fmt.Println(ss)
	fmt.Printf("%s%s\n", s3, s4)
	//用Sprintf可以将拼接结果返回
	ss1 := fmt.Sprintf("%s%s", s3, s4)
	fmt.Println(ss1)

	//字符串的分割
	ret := strings.Split(s3, ".")
	fmt.Println(ret)

	//判断是否包含
	//返回bool类型值
	fmt.Println(strings.Contains(s4, "理性"))
	fmt.Println(strings.Contains(s4, "理想"))

	//前缀判断
	fmt.Println(strings.HasPrefix(s4, "理想"))
	//后缀判断
	fmt.Println(strings.HasSuffix(s4, "理想"))

	//位置从0开始算
	//子串出现的位置
	ss2 := "abcdeb"
	fmt.Println(strings.Index(ss2, "c"))
	//子串最后一次出现的位置
	fmt.Println(strings.LastIndex(ss2, "b"))

	//拼接
	fmt.Println(strings.Join(ret, "+"))
}

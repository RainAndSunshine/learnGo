package main

import (
	"fmt"
	"strconv"
)

//strconv标准库
//实现不同类型之间的转换

func main() {
	//从字符串中解析出对应的数据
	str := "10000"
	//参数：字符串   进制（十进制）   位数（int64）(如果传的是0，自动返回int类型）
	//返回值是int64类型的，这样保证数据转换过程中，位数不会丢失，之后可以直接将int64强转成需要的类型
	ret1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Printf("parse int err:%v\n", err)
		return
	}
	fmt.Printf("type:%T,value:%v\n", ret1, ret1)

	//如果需要直接转换成int类型，有个现成的方法
	retInt, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("parse int err:%v\n", err)
		return
	}
	fmt.Printf("type:%T,value:%v\n", retInt, retInt)

	//从字符串中解析出bool值
	boolStr := "true"
	retBool, err := strconv.ParseBool(boolStr)
	if err != nil {
		fmt.Printf("parse bool err:%v\n", err)
		return
	}
	fmt.Printf("type:%T,value:%v\n", retBool, retBool)

	//从字符串中解析出浮点数
	floatStr := "3.1415926"
	retFloat, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Printf("parse float err:%v\n", err)
		return
	}
	fmt.Printf("type:%T,value:%v\n", retFloat, retFloat)

	//把数字类型转换为字符串类型
	a := 10000
	ret2 := fmt.Sprintf("%d", a)
	fmt.Printf("type:%T,value:%v\n", ret2, ret2)

	//将int类型转换为字符串类型
	retString := strconv.Itoa(a)
	fmt.Printf("type:%T,value:%v\n", retString, retString)
}

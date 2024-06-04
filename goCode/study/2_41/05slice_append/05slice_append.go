package main

import "fmt"

func main() {
	//append() 为切片追加元素
	//append追加元素，原来的底层数组放不下的时候，Go底层就会把底层数组换一个
	//必须用变量接受append的返回值
	s1 := []string{"北京", "上海", "深圳"}
	fmt.Printf("s1:%v  len(s1):%d  cap(s1):%d\n", s1, len(s1), cap(s1))
	//调用append函数必须用原来的切片变量接受返回值
	s1 = append(s1, "广州")
	fmt.Printf("s1:%v  len(s1):%d  cap(s1):%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "杭州", "苏州")
	fmt.Printf("s1:%v  len(s1):%d  cap(s1):%d\n", s1, len(s1), cap(s1))
	sss := []string{"西安", "成都", "重庆"}
	s1 = append(s1, sss...) //...表示拆开
	fmt.Printf("s1:%v  len(s1):%d  cap(s1):%d\n", s1, len(s1), cap(s1))

	//切片的扩容策略
	//首先判断，如果新申请的容量大于2倍的旧容量，最终容量就是新申请的容量
	//否则判断，如果旧切片的长度小于1024，则最终容量就是旧容量的两倍
	//否则判断，如果旧切片的长度大于等于1024，则最终容量从旧容量开始循环增加原来的1/4，直到最终容量大于等于新申请的容量
	//如果最终容量计算值溢出，则最终容量就是新申请的容量

	//切片扩容还会根据切片中元素的类型不同而做出不同的处理
}

package main

import (
	"fmt"
	"strings"
)

// 统计一个字符串中每个单词出现的次数。比如：“how do you do”中how=1、do=2、you=1
// 思路：
// 利用字符串切割，将句子转换成多个单词
// 利用切片存放单词
// 利用map的特性(成对出现)，将单词和对应次数联系在一起
func main() {
	//创建一个字符串
	s1 := "can you can a can as a caner can can a can"
	//创建切片
	s2 := make([]string, 15, 20)
	//将字符串按空格切割，放入切片内
	s2 = strings.Split(s1, " ")
	//创建map
	m1 := make(map[string]int, 15)
	//遍历切片，将每个单词与对应次数添加至map中
	for _, v := range s2 {
		//1.如果map不存在v这个key，初始化次数为1
		//2.如果map存在v这个key，次数+1
		if _, ok := m1[v]; !ok {
			m1[v] = 1
		} else {
			m1[v]++
		}
	}
	//遍历map，输出结果
	for key, value := range m1 {
		fmt.Println(key, value)
	}
}

package split_string

import "strings"

//实现字符串切割
//不使用内置的split函数

func Split(str, sep string) []string {
	var res []string
	//获取要切割的索引
	index := strings.Index(str, sep)
	//循环追加，索引前后的元素
	for index >= 0 {
		res = append(res, str[:index])
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}
	res = append(res, str)
	return res
}

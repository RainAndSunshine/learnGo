package main

import "fmt"

var s11 = [...]string{"abba", "trend", "yes", "bccb"}

func main() {
	//回文判断，判断字符串是否回文
	for i := 0; i < len(s11); i++ {
		for j := 0; j <= len(s11[i]); j++ {
			if s11[i][j] != s11[i][len(s11[i])-1-j] {
				fmt.Println("第", i+1, "个字符串不是回文的")
				break
			}
			if j >= len(s11[i])/2 {
				fmt.Println("第", i+1, "个字符串是回文的")
				break
			}
		}
	}

	//判断中文的，需要用切片，元素类型为[]rune
	ss := "a山西运煤车煤运山西a"
	r := make([]rune, 0, len(ss))
	for _, v := range ss {
		r = append(r, v)
	}
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
}

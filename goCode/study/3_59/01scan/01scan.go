package main

import "fmt"

func main() {
	//获取用户输入
	var s string
	//记住要传指针
	fmt.Scan(&s)
	fmt.Println("输入的内容是：", s)

	var (
		name  string
		age   int
		class string
	)
	fmt.Scanf("%s %d %s\n", &name, &age, &class)
	fmt.Printf("姓名：%s  年龄：%d  班级：%s\n", name, age, class)

	//跟Scanf的区别是，它可以自动扫描到换行符，不用打出来
	fmt.Scanln(&name, &age, &class)
	fmt.Printf("姓名：%s  年龄：%d  班级：%s\n", name, age, class)

}

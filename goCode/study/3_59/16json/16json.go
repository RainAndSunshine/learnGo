package main

import (
	"encoding/json"
	"fmt"
)

//结构体与json
//1.序列化：	把Go语言中的结构体变量 --> json格式的字符串
//2.反序列化	把json格式的字符串	   --> Go语言中能够识别的结构体变量

//1.架构体内部的字段首字母要大写！！！不大写别人是访问不到的
//2.反序列化时要传递指针！

type person5 struct {
	//字段小写的话，外部的包就无法调用
	Name string `json:"name" db:"name" ini:"name"` //空格+反引号+包+“名称”
	Age  int    `json:"age"`                       //可以保证序列化之后是小写
}

func main() {
	d1 := person5{
		Name: "周林",
		Age:  9000,
	}
	//序列化
	b, err := json.Marshal(d1)
	if err != nil {
		fmt.Println("marshal failed，err:%v:", err)
		return
	}
	fmt.Printf("%#v\n", string(b))
	//反序列化
	str := `{"name":"理想","age":9000}`
	var d2 person5
	json.Unmarshal([]byte(str), &d2) //传指针是为了能在json.Unmarshal内部修改d2的值
	fmt.Printf("%#v\n", d2)
}

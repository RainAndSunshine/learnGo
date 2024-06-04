package main

import (
	"encoding/json"
	"fmt"
)

// json
type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name":"zhoulin","age":9000}`
	var p person
	//第二个参数是空指针，什么类型都能传进去
	//但编译器是怎么判断具体的类型呢
	//这就利用了反射原理
	json.Unmarshal([]byte(str), &p)
	fmt.Println(p.Name)
	fmt.Println(p.Age)
}

package main

import "fmt"

var (
	n = 5
)

func main() {
	//switch
	switch n {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的数字")
	}

	//一个分支可以有多个值，多个case值中间用英文逗号分隔
	age := 10
	switch age {
	case 1, 2, 3, 4, 5:
		fmt.Println("1~5")
	case 6, 7, 8, 9, 10:
		fmt.Println("6~10")
	default:
		fmt.Println("数字不在1~10之间")
	}

	//switch前可以有执行语句
	switch pp := 18; pp {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
		fmt.Println("1~18no_name_func")
	case 11, 12, 13, 14, 15, 16, 17:
		fmt.Println("11~17func_type")
	case 18:
		fmt.Println("刚满十八岁")
	default:
		fmt.Println("不在1~18之间")
	}

	//分支可以使用表达式，这时候switch后不需要再加变量
	i := 10
	//switch后不需要跟变量名
	switch {
	case i%2 == 0:
		fmt.Println("偶数")
	case i%2 != 0:
		fmt.Println("奇数")
	}

	//fallthrough和goto都别写，容易出问题
	//fallthrough可以执行满足条件的case的下一个case
	s := "a"
	switch s {
	case "a":
		fmt.Println("a")
		fallthrough //下一个case也可以执行
	case "b":
		fmt.Println("b")
		fallthrough //下一个case也可以执行
	case "c":
		fmt.Println("c")
		//没有fallthrough，下一个case不执行了
	case "d":
		fmt.Println("d")
	default:
		fmt.Println("都不是")
	}

	//goto语句，直接跳转
	for i = 0; i < 10; i++ {
		fmt.Print(i)
		if i == 4 {
			goto label
		}
	}
	fmt.Println("这句话被跳过了我去")
label:
	fmt.Println()
	fmt.Println("我在这里")
}

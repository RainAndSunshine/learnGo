package main

import (
	"fmt"
	"path"
	"runtime"
)

func f1() {
	//参数是层数，如果是0，返回的line就是这条语句所在的行数,找到f1里面调用函数的地方
	//如果是1，则去上一个调用这个语句的地方,找到f2里面调用的地方
	//如果是2，则再往上找一层,找到main函数里面调用的地方
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Println("runtime.Caller() failed\n")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(file)
	//找到file的路径
	fmt.Println(path.Base(file))
	fmt.Println(line)
}

func f2() {
	f1()
}

func main() {
	f2()
}

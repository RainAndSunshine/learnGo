package main

import (
	"fmt"
	"time"
)

//goroutine

func hello(i int) {
	fmt.Println("hello", i)
}

// 程序启动之后会创建一个main goroutine
func main() {
	for i := 0; i <= 1000; i++ {
		//go hello(i)
		//启动goroutine是需要时间的
		go func() {
			fmt.Println(i)
		}()
	}
	fmt.Println("main")
	time.Sleep(1 * time.Second)
	//main函数结束了 由main函数启动的goroutine也都结束了
}

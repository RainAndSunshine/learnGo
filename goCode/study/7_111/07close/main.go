package main

import "fmt"

//关闭通道
//关闭通道之后，取完值后，通道为空就不可以取值了,取到的值是0值，ok是false

//通道一般不需要手动关闭
//channel是一种类型，在程序运行中，是一直在用的
//在程序结束后，会自动进行垃圾回收，自动关闭
//但是如果有range通道的情况
//不关闭通道，将会一直取值，造成阻塞

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 20
	close(ch1)
	//for x := range ch1 {
	//	fmt.Println(x)
	//}
	x0, ok0 := <-ch1
	fmt.Println(x0, ok0)
	x1, ok1 := <-ch1
	fmt.Println(x1, ok1)
	//读完数据后，返回类型的零值
	x, ok := <-ch1
	fmt.Println(x, ok)
}

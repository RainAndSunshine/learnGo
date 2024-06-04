package main

import (
	"fmt"
	"sync"
)

//channel练习
//1.启动一个goroutine，生成100个数发送到ch1
//2.启动一个goroutine，从ch1中取值，计算其平方放到ch2中
//3.在main函数中，从ch2中取值打印出来

var ch1 = make(chan int, 100)
var ch2 = make(chan int, 100)
var wg sync.WaitGroup

var once sync.Once

func f1() {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		ch1 <- i
	}
	//关闭之后不可以写，但可以读
	close(ch1)
}

func f2() {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	//参数是匿名函数
	once.Do(func() { close(ch2) }) //确保某个函数只执行一次
}

func main() {
	wg.Add(3)
	go f1()
	go f2()
	go f2() //开启两个goroutine同步进行
	wg.Wait()
	for x := range ch2 {
		fmt.Println(x)
	}
}

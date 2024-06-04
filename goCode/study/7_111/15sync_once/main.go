package main

//当某个操作必须确保只能执行一次的时候，使用sync.Once

import (
	"fmt"
	"sync"
)

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
	//参数是匿名函数，使用闭包
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

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//goroutine什么时候结束
//goroutine对应的函数结束了，goroutine结束了
//main函数执行完了，由main函数创建的那些goroutine都结束了

//waitGroup

func f() {
	//现在随机数种子已经被弃用了，rand自带了
	//seed的参数是int64类型
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		//现在的rand自带了seed，每一次执行都不一样
		ret1 := rand.Int()
		ret2 := rand.Intn(10) //0<=x<10
		fmt.Println(ret1, ret2)
		//取负数
		fmt.Println(0-ret1, 0-ret2)
	}
}

func f1(i int) {
	//函数结束后，wg-1
	defer wg.Done()
	//sleep里面接受的是time.Duration类型，随机数是int类型
	//不同类型不能进行运算，所以需要强制类型转换
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	fmt.Println(i)
}

// 创建一个waitGroup类型变量
var wg = sync.WaitGroup{}

func main() {
	//f()
	//wg.Add(10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}
	//如何知道这10个goroutine都结束了，然后结束main
	wg.Wait() //等待wg的计数器减为0
}

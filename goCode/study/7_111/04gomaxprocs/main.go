package main

import (
	"fmt"
	"runtime"
	"sync"
)

// GOMAXPROCS
//一个软件一个进程，一个进程可以创建多个线程去干活

//Go语言中的 操作系统线程 和 goroutine 之间的关系
//1.一个操作系统线程对应用户态多个goroutine
//2.go程序可以同时使用多个操作系统线程
//3.goroutine和OS线程是多对多的关系，即 m:n

//goroutine调度模型  GMP
//M:N 把M个goroutine分配给N个操作系统线程去执行
//goroutine初始栈大小为2KB，是用户态的线程，比操作系统线程更轻量级

var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A: %d\n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B: %d\n", i)
	}
}

func main() {
	//占用的cpu数量
	runtime.GOMAXPROCS(1) //默认cpu的逻辑核心数，默认跑满整个cpu
	fmt.Println(runtime.NumCPU())
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}

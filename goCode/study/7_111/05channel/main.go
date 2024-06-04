package main

import (
	"fmt"
	"sync"
)

//channel

//Go语言的并发模型是CSP 提倡 通过通信共享内存 而不是通过共享内存而实现通信
//如果说goroutine是Go程序并发的执行体，channel就是它们之间的连接
//channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制
//Go语言中的通道（channel）是一种特殊的类型
//通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序
//每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型

//var 通道名 chan 通道中元素的类型
//chan是一个引用类型，必须初始化，使用make函数

//一般不存放数组，切片，map，结构体等比较大的类型
//一般存放他们的指针

//通道的操作
//发送	ch1 <- 999

//接收	x := <- ch1	(没变量接收 数据就扔了)

//关闭	close()

//我们可以使用内置的len函数获取通道内元素的数量，使用cap函数获取通道的容量，虽然我们很少会这么做。

var b chan int

var wg sync.WaitGroup

func noBufChannel() {
	fmt.Println(b) //nil
	wg.Add(1)
	b = make(chan int) //不带缓冲区的通道初始化
	//无缓冲的通道必须有至少一个接收方才能发送成功
	go func() { //启动另外一个goroutine去等待，直到通道b中有值发送，然后让x接收b
		defer wg.Done()
		x := <-b
		fmt.Println("后台goroutine从通道b中取到了", x)
	}()
	b <- 10 //通道b没缓冲区，放不进去，无法进行下一步操作
	// 只能直等待数据接收，如果不从通道中发送出去，就会卡死main函数的goroutine
	fmt.Println("10发送到通道b中了")
	wg.Wait()
	close(b)
}

func BufChannel() {
	fmt.Println(b)
	//通道的值尽量小一点，如果太大，就传地址
	b = make(chan int, 1) //带缓冲区的通道初始化，容量为1
	b <- 10
	fmt.Println("10发送到通道b中了")
	b <- 20 //如果多传了，就发送不了，程序卡住了，死锁了
	x := <-b
	fmt.Println("从通道b中取到了", x)
	close(b) //可以不关闭通道，通道是可以自动回收的
}

func main() {
	//noBufChannel()
	BufChannel()
}

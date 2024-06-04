package main

import "fmt"

//Select 的使用方式类似于之前学到的 switch 语句 它也有一系列 case 分支和一个默认的分支
//每个 case 分支会对应一个通道的通信（接收或发送）过程
//select 会一直等待，直到其中的某个 case 的通信操作完成时，就会执行该 case 分支对应的语句

func main() {
	//创建一个缓冲区为1的通道
	ch := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		//select
		//当奇数的时候，通道没有值，无法发送出来，只能接收值
		//当偶数的时候，通道有值且缓冲区已满，无法接收值，只能发送值出来
		//从而实现输出奇数
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}

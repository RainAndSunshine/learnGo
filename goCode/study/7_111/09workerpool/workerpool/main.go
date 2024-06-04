package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(id int, jobs <-chan int, results chan<- int) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//三个线程，并发
	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)
	//开始等待
	//等到三个线程都结束，说明数据已全部发送到results通道中
	wg.Wait()
	//关闭通道，如果不关闭通道，range读完数据将不会停止，导致所有线程（这里就剩下一个main线程）阻塞，导致死锁
	close(results)
	for x := range results {
		fmt.Println(x)
	}
}

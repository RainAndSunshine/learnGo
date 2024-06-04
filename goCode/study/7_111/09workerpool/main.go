package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
使用goroutine和channel实现一个计算int64类型随机数各位数和的程序
1. 开启一个goroutine循环生成int64类型的随机数，发送到jobChan
2. 开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
3. 主goroutine从resultChan取出结果并打印到终端输出
*/

type job struct {
	j int64
}

type result struct {
	job job
	sum int64
}

var wg sync.WaitGroup
var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)

func f1(jobChan chan<- *job) {
	defer wg.Done()
	for {
		x := rand.Int63()
		newJob := &job{
			j: x,
		}
		jobChan <- newJob
		time.Sleep(time.Second / 2)
	}
}

func f2(jobChan <-chan *job, resultChan chan<- *result) {
	defer wg.Done()
	for {
		x := <-jobChan
		sum := int64(0)
		var tmp int64 = x.j
		for tmp > 0 {
			sum += tmp % 10
			tmp /= 10
		}
		newResult := &result{
			job: *x,
			sum: sum,
		}
		resultChan <- newResult
	}
}

func main() {
	wg.Add(1)
	go f1(jobChan)
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go f2(jobChan, resultChan)
	}
	for res := range resultChan {
		fmt.Printf("job:%d sum:%d \n", res.job.j, res.sum)
	}
	wg.Wait()
}

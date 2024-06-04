package main

import (
	"fmt"
	"sync"
)

//锁

//互斥锁
//互斥锁是一种常用的控制共享资源访问的方法，它能够保证同一时间只有一个goroutine可以访问共享资源
//避免 goroutine 共同修改数据造成数据错误

var x = 0
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	defer wg.Done()
	for i := 0; i < 50000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

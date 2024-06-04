package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

//原子操作
//针对整数数据类型（int32、uint32、int64、uint64）我们还可以使用原子操作来保证并发安全，通常直接使用原子操作比使用锁操作效率更高

var x int64
var lock sync.Mutex
var wg sync.WaitGroup

func add() {
	defer wg.Done()
	lock.Lock()
	x++
	lock.Unlock()
}

func addByAtomic() {
	defer wg.Done()
	//传的是x的指针，因为函数传参传的是拷贝
	atomic.AddInt64(&x, 1)
}

func main() {
	start1 := time.Now()
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Println(x)
	fmt.Println(time.Since(start1))

	start2 := time.Now()
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go addByAtomic()
	}
	wg.Wait()
	fmt.Println(x)
	fmt.Println(time.Since(start2))
}

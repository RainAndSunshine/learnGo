package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写互斥锁——运用在读远大于写的时候
// 当一个goroutine获取读锁的时候，其他goroutine还可以读
// 当一个goroutine获取写锁的时候，其他goroutine无论获取的是读锁还是写锁，都要等待

var (
	x      = 0
	lock   sync.Mutex
	rwLock sync.RWMutex
	wg     sync.WaitGroup
)

func read() {
	defer wg.Done()
	//lock.Lock()
	rwLock.RLock()
	time.Sleep(time.Millisecond)
	//lock.Unlock()
	rwLock.RUnlock()
}

func write() {
	defer wg.Done()
	//lock.Lock()
	rwLock.Lock()
	time.Sleep(time.Millisecond * 5)
	//lock.Unlock()
	rwLock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	//这个相当于 time.Now().Sub(start)
	fmt.Println(time.Now().Sub(start))
	fmt.Println(time.Since(start))
}

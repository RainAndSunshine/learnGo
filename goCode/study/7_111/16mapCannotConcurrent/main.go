package main

//Go语言中内置的map不是并发安全的

import (
	"fmt"
	"strconv"
	"sync"
)

//var lock sync.Mutex

var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func main() {
	wg := sync.WaitGroup{}
	//并发操作map会造成 concurrent map writes 错误
	//Go语言内置的map是不支持并发操作的
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			//加锁可以解决这个问题，但是效率较低
			//lock.Lock()
			set(key, n)
			//lock.Unlock()
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

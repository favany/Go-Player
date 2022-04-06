package main

import (
	"fmt"
	"sync"
)

var (
	x    int64
	wg   sync.WaitGroup
	lock sync.Mutex // 互斥锁
)

func accumulate() {
	for i := 0; i < 100000; i++ {
		lock.Lock() // 加锁
		x = x + 1
		lock.Unlock() // 释放锁
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go accumulate()
	go accumulate()
	wg.Wait()
	fmt.Println(x)
}

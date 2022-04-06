package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写互斥锁

var (
	x      int64
	wg     sync.WaitGroup
	rwlock sync.RWMutex
)

func read() {
	rwlock.RLock()
	time.Sleep(time.Millisecond) // 模拟读锁
	rwlock.RUnlock()
	wg.Done()
}

func write() {
	rwlock.Lock()
	time.Sleep(time.Millisecond * 10) // 模拟写锁
	rwlock.Unlock()
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 6666; i++ {
		wg.Add(1)
		go read()
	}

	for i := 0; i < 66; i++ {
		wg.Add(1)
		go write()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

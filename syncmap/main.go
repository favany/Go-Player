package main

import (
	"fmt"
	"sync"
)

var (
	wg     sync.WaitGroup
	rwlock sync.RWMutex
)

var m = make(map[int]int)

func get(key int) int {
	rwlock.RLock()
	ret := m[key]
	rwlock.RUnlock()
	return ret
}

func set(key int, value int) {
	rwlock.Lock()
	m[key] = value
	rwlock.Unlock()
}

func main() {
	for i := 0; i < 666; i++ {
		wg.Add(1)
		go func(i int) {
			set(i, i+100)                              // 设置 map 键值对
			fmt.Printf("key:%v value:%v  ", i, get(i)) // 打印键值对
			wg.Done()
		}(i)
	}
	wg.Wait()
}

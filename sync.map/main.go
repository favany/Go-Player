package main

import (
	"fmt"
	"sync"
)

// sync.Map 并发安全的map

var (
	wg sync.WaitGroup
	m2 sync.Map
)

func main() {
	for i := 0; i < 66; i++ {
		wg.Add(1)
		go func(i int) {
			m2.Store(i, i+100) // 设置 map 键值对
			value, _ := m2.Load(i)
			fmt.Printf("key:%v value:%v", i, value) // 打印键值对
			wg.Done()
		}(i)
	}
	wg.Wait()
}

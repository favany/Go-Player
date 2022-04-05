package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(10000) // 计数器加一

	for i := 0; i < 10000; i++ {
		go func(i int) {
			fmt.Println("hello", i) // 很多的 i 打印了相同的值
			wg.Done()
		}(i)
	}
	fmt.Println("main")
	wg.Wait()

}

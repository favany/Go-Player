package main

import (
	"fmt"
	"sync"
)

// goroutine1

var wg sync.WaitGroup

func hello() {
	fmt.Println("Hello, Gophist!")
	wg.Done() // 通知 wg 把计数器-1
}

func main() { // 开启一个主 goroutine1 去执行 main 函数
	wg.Add(10000) // 技术牌 + 1
	for i := 0; i < 10000; i++ {
		go hello() // 开启了一个 goroutine1 去执行 hello 这个函数
	}
	fmt.Println("Hello, main!")
	wg.Wait() // 阻塞，等所有小弟都干完活后，才收兵
}

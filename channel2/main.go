package main

import "fmt"

/*
两个 goroutine, 两个 channel
1. 生成 0 - 100 的数字发送到 ch1
2. 从 ch1 取出数据计算它的平方，把结果发送到 ch2 中
*/

// 生成 0 ~ 100 数字发送到 ch1
//func f1(ch chan<- int) {
func f1(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

//func f2(ch1 <-chan int, ch2 chan<- int) {
func f2(ch1 chan int, ch2 chan int) {
	// 从通道中取值的方式 1
	for {
		tmp, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- tmp * tmp
	}
	close(ch2)

}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 200)

	go f1(ch1)
	go f2(ch1, ch2)

	// 从通道中取值的方式2
	for ret := range ch2 {
		fmt.Println(ret)
	}

}

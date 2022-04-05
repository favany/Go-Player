package main

/*
两个 goroutine, 两个 channel
1. 生成 0 - 100 的数字发送到 ch1
2. 从 ch1 取出数据计算它的平方，把结果发送到 ch2 中
*/

// 生成 0 ~ 100 数字发送到 ch1
func f1(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- 1
	}
	close(ch)
}

func f2(ch1 chan int, chan2 chan int) {
	for {
		tmp, ok := <-ch1
		if !ok {
			break
		}
		chan2 <- tmp * tmp

	}

}

func main() {

}

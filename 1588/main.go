package main

import "fmt"

func sumOddLengthSubarrays(arr []int) int {
	// 获取数组的长度
	len := len(arr)
	// 求和的值
	sum := 0
	// 特殊情况：空
	if len == 0 {
		return 0
	}
	// 数组长度循环: 1 , 3, 5 ...
	for l := 1; l <= len; l += 2 {
		//fmt.Println("l", l)
		// 数组从哪里开始？: 0, 1, 2, ...
		for startIndex := 0; startIndex <= len-l; startIndex++ {
			//fmt.Println("startIndex", startIndex)
			// 数组内循环
			for i := 0; i < l; i++ {
				//fmt.Println("i", i, arr[startIndex+i])
				sum += arr[startIndex+i]
			}
		}
	}
	return sum
}

func main() {
	arr := []int{1, 4, 2, 5, 3}
	re := sumOddLengthSubarrays(arr)
	fmt.Println(re)
}

package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	numStruct := make(map[int]int)
	// 一共有几对
	count := 0
	for _, num := range nums {
		numStruct[num] += 1
	}

	fmt.Println(numStruct)
	for _, val := range numStruct {
		if val > 1 {
			count += factorial(val) / 2
		}
	}

	return count
}

func factorial(num int) (re int) {
	re = 1
	for a := num; a >= 1; a-- {
		fmt.Println(a)
		re *= a
	}
	fmt.Println("阶乘", re)
	return
}

func main() {
	numIdenticalPairs([]int{1, 1, 1, 1})
}

package main

import "fmt"

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	// 1. 先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 // 黑棋
	chessMap[2][3] = 2 // 蓝棋

	// 2. 输出原始数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	// 3. 转成稀疏数组
	// 思路
	// （1）遍历 chessMap，如果发现有一个元素的值不等于0，创建一个 ValNode 结构体
	// （2）将其放入到对应的切片即可。

	var sparseArr []ValNode

	// 标准的一个稀疏数组应该先有一个记录元素的二维数组的规模（行和列、默认值）

	// 首先，创建一个 ValNode 的值节点，添加到 sparseArr
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)

	// 其次，遍历 chessMap ， 把非 0 值添加到 sparseArr
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				// 创建一个ValNode 值节点
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	// 输出稀疏数组
	fmt.Println("当前的稀疏数组是：")
	for i, valNode := range sparseArr {
		fmt.Printf("%d： %d %d %d", i, valNode.row, valNode.col, valNode.val)
		fmt.Println()
	}
}

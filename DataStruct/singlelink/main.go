package main

import "fmt"

// 单链表

// 定义一个HeroNode

type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode // 这个表示指向下一个节点
}

// 给链表插入一个结点

// 编写第一种插入方式，在单链表的最后加入.

func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	// 思路
	// 1. 创建一个辅助结点（跑龙套，帮忙）
	temp := head
	// 2. 先找到该链表的最后这个结点
	for {
		if temp.next == nil { // 表示找到最后
			break
		}
		temp = temp.next // 让temp指向下一个结点，实现循环♻️
	}
	// 3. 把原先最后一个结点的 next 指向 newHeroNode，即将 newHeroNode 加入到链表的最后
	temp.next = newHeroNode
	fmt.Println(temp)

}

// 显示链表的所有结点信息
// 韩顺平老师的写法
func ListHeroNode(head *HeroNode) {

	// 1. 创建一个辅助结点[跑龙套，帮忙]，从第一个结点开始
	temp := head

	// 2. 先判断该链表是不是一个空链表
	if temp.next == nil {
		fmt.Println("该链表空空如也...")
		return
	}

	// 2. 遍历这个链表
	for {
		fmt.Printf("[%d, %s, %s] -> \n",
			temp.next.no,
			temp.next.name,
			temp.next.nickname,
		)
		// 下一个结点
		temp = temp.next
		// 判断是否存在该结点
		if temp.next == nil {
			break
		}
	}

}

// 我的写法

func ListHeroNode2(head *HeroNode) {

	// 1. 创建一个辅助结点[跑龙套，帮忙]，从第一个结点开始
	temp := head.next

	// 2. 先判断该链表是不是一个空链表
	if temp == nil {
		fmt.Println("该链表空空如也...")
		return
	}

	// 2. 遍历这个链表
	for {
		fmt.Printf("[%d, %s, %s] -> \n",
			temp.no,
			temp.name,
			temp.nickname,
		)
		// 下一个结点
		temp = temp.next
		// 判断是否存在该结点
		if temp == nil {
			break
		}
	}

}

func main() {
	// 1. 先创建一个头节点
	head := &HeroNode{}

	// 2. 创建一个新的 HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}

	// 3. 加入
	InsertHeroNode(head, hero1)

	// 4. 显示
	fmt.Println("韩顺平老师的写法")
	ListHeroNode(head)
	fmt.Println("我的写法")
	ListHeroNode2(head)
}

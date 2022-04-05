package main

import "fmt"

// 单链表

// 定义一个Node

type Node struct {
	no   int
	name string
	next *Node // 这个表示指向下一个节点
}

// 给链表插入一个结点

// 编写第一种插入方式，在单链表的最后加入

func InsertHeroNode(head *Node, newNode *Node) {
	// 1. 创建一个flag 🚩结点，用于遍历链表中的每个节点
	flag := head
	// 2. 通过遍历，找到该链表的最后一个结点
	for {
		if flag.next == nil { // 表示找到最后
			break
		}
		flag = flag.next // 让temp指向下一个结点，实现循环♻️
	}
	// 3. 把原先最后一个结点的 next 指向 newNode，即将 newNode 加入到链表的最后
	flag.next = newNode

}

// 显示链表的所有节点信息

func ListNode(head *Node) {

	// 1. 创建一个辅助结点作为迭代的flag 🚩，从第一个结点开始
	flag := head.next

	// 2. 先判断该链表是不是一个空链表
	if flag == nil {
		fmt.Println("该链表空空如也...")
		return
	}

	// 2. 遍历这个链表
	for {
		fmt.Printf("[%d, %s] ",
			flag.no,
			flag.name,
		)
		// 下一个结点
		flag = flag.next
		// 判断是否存在该结点
		if flag == nil {
			break
		} else {
			fmt.Print("-> ")
		}
	}

	fmt.Println()
}

// 删除链表中的一个节点

func DeleteNode(head *Node, deleteNode *Node) {
	// 1. 创建一个flag 🚩结点，用于遍历链表中的每个节点
	flag := head
	// flag.next 从第一个节点开始遍历
	for {
		// deleteNode 和 flag.next 它们指向的下一个节点相同，代表它们本身相同。
		// 要删除 flag.next，就是要把 flag 指向 flag.next.next
		if deleteNode.next == flag.next.next {
			flag.next = flag.next.next
			return
		}

		if flag.next == nil {
			return
		} else {
			flag = flag.next
		}
	}

}

func main() {
	// 1. 先创建一个头节点
	head := &Node{}

	// 2. 创建新的节点 Node
	node1 := &Node{
		no:   1,
		name: "宋江",
	}

	node2 := &Node{
		no:   2,
		name: "卢俊义",
	}

	node3 := &Node{
		no:   3,
		name: "林冲",
	}

	// 3. 加入
	InsertHeroNode(head, node1)
	InsertHeroNode(head, node2)
	InsertHeroNode(head, node3)

	// 4. 显示
	ListNode(head)

	// 5. 删除
	DeleteNode(head, node2)

	// 6. 显示删除后的结果
	ListNode(head)

}

package main

import (
	"errors"
	"fmt"
	"os"
)

// 使用数组实现队列的思路
// 1. 创建一个数组 array，是作为队列的一个字段
// 2. front 初始化为 -1
// 3. rear 表示队列尾部，初始化为 -1
// 4. 完成队列的基本查找

// 加入数据到队列 AddQueue()
// 从队列取出数据 GetQueue()
// 显示队列      ShowQueue()

// 使用一个结构体管理队列

type Queue struct {
	maxSize int
	array   [5]int // 数组 => 模拟队列
	front   int    // 表示指向队列首
	rear    int    // 表示指向队列的尾部
}

// 添加数据到队列

func (q *Queue) AddQueue(val int) (err error) {
	// 先判断队列是否已满
	if q.rear == q.maxSize-1 { // 重要‼️提示：rear 是队列尾部（含最后元素）
		return errors.New("队列已满，无法加入！")
	}

	q.rear++ // rear 后移
	q.array[q.rear] = val
	return
}

// 显示队列

func (q *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是：")
	// q.front 不含队首的元素
	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, q.array[i])
	}
	fmt.Println()
	fmt.Println()
}

// 从队列中取出数据

func (q *Queue) GetQueue() (val int, err error) {
	// 先判断队列是否为空
	if q.rear == q.front { // 队空
		return -1, errors.New("队列是空")
	}
	q.front++

	return val, err
}

func main() {
	// 先创建一个队列
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}

	var key string
	var val int
	for {
		fmt.Println("1. 输入 add 表示添加数据到队列")
		fmt.Println("2. 输入 get 表示从队列获取数据（尚未完工）")
		fmt.Println("3. 输入 show 表示显示队列")
		fmt.Println("4. 输入 exit 表示退出队列")
		fmt.Println("请输入 add / get / show / exit ：")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入要加入队列的数值：")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("欧耶，加入队列成功!!!")
				fmt.Println()
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("队列中的数为", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}

// 小结：
// 对上面代码的小结和说明：
// 上面代码实现了基本队列结构，但是没有有效地利用数组空间。
// 请思考，如何使用数组实现一个环形的队列。

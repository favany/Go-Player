package main

import (
	"errors"
	"fmt"
	"os"
)

// 数组模拟环形队列
// 对前面的数组模拟队列的优化，充分利用数组。因此将数组看作是一个环形的，通过取模的方式来实现即可。
// 提醒：
// 1. 尾索引的下一个为头索引时，表示队列满，即将队列容量空出一个作为约定，这个在做队列满的时候需要注意
// (tail + 1) % maxSize == head 满
// 2. tail == head 空

// CircleQueue 使用一个结构体管理环形队列
type CircleQueue struct {
	maxSize int    // 4
	array   [5]int // 数组
	head    int    // 指向队首
	tail    int    // 指向队尾
}

// 入队列 push

func (c *CircleQueue) Push(val int) (err error) {
	if c.IsFull() {
		return errors.New("队列已经满了")
	}
	// tail 在队列的尾部，但是不包含最后的元素
	c.array[c.tail] = val // 把值给尾部
	c.tail = (c.tail + 1) % c.maxSize
	return
}

// 弹出（队列最后一个元素） pop

func (c *CircleQueue) Pop() (val int, err error) {
	if c.IsEmpty() {
		return 0, errors.New("队列为空！")
	}
	// 取出, head 是指队首，并且包含队首元素
	val = c.array[c.head]
	c.head = (c.head + 1) % c.maxSize
	return
}

// 显示队列

func (c *CircleQueue) ListQueue() {
	fmt.Println("环形队列的情况如下：")
	// 取出当前队列有多少个元素
	size := c.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}

	// 设计一个辅助的变量，指向head
	tempHead := c.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, c.array[tempHead])
		tempHead = (tempHead + 1) % c.maxSize
	}
	fmt.Println()
}

// 判断环形队列为满

func (c *CircleQueue) IsFull() bool {
	return (c.tail+1)%c.maxSize == c.head
}

// 判断环形队列为空

func (c *CircleQueue) IsEmpty() bool {
	return c.tail == c.head
}

// 取出环形队列中有多少个元素

func (c *CircleQueue) Size() int {
	// 这是关键
	//return (c.tail + c.maxSize - c.head) % c.maxSize
	return (c.tail - c.head) % c.maxSize
}

func main() {
	// 初始化一个环形队列
	queue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
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
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("欧耶，加入队列成功!!!")
				fmt.Println()
			}
		case "get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("队列中的数为", val)
			}
		case "show":
			queue.ListQueue()
		case "exit":
			os.Exit(0)
		}
	}

}

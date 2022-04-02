package main

import "fmt"

// 使用 值接收者实现接口 和 使用 指针接受者实现接口 的区别

type mover interface {
	move()
}

type person struct {
	name string
	age  int8
}

func (p person) move() {
	fmt.Printf("%s在跑...", p.name)
}

func main() {
	var m1 mover
	var m2 mover
	p1 := person{ // person 类型的值
		name: "Gophist",
		age:  18,
	}
	p2 := &person{ // person 类型指针
		name: "Bingo",
		age:  22,
	}
	m1 = p1
	m2 = p2
	m1.move()
	m2.move()
	fmt.Println(m1, m2)
}

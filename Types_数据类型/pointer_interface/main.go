package main

import "fmt"

// 使用 值接收者实现接口 和 使用 指针接受者实现接口 的区别

// 接口的嵌套
type animal interface {
	mover
	sayer
}

type mover interface {
	move()
}

type sayer interface {
	say()
}

type person struct {
	name string
	age  int8
}

// 使用值接收者实现接口：类型的值和类型的指针都能够保存到结构的变量中

func (p person) move() {
	fmt.Printf("%s在跑...", p.name)
}

// 使用指针接收者实现接口：只有类型指针能够保存到接口变量中
func (p *person) move2() {
	fmt.Printf("%s在跑...", p.name)
}

func (p person) say() {
	fmt.Printf("%s在说...", p.name)
}

func main() {
	var (
		m1 mover
		m2 mover
	)

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

	var ani animal
	ani = &person{
		name: "Gopher",
		age:  3,
	}
	ani.say()
	ani.move()
	fmt.Println(ani)
}

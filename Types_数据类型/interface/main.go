package main

import (
	"fmt"
)

// 为什么需要接口？
type dog struct {
}

func (d dog) say() {
	fmt.Println("汪汪汪~")
}

type cat struct {
}

func (c cat) say() {
	fmt.Println("喵喵喵~")
}

type person struct {
	name string
}

func (p person) say() {
	fmt.Println("啊啊啊~")
}

// 接口不管具体是什么类型，它只管实现什么方法
// 定义一个类型，一个抽象的类型，只要实现了 say() 这个方法的类型都可以称为 sayer() 类型
type sayer interface {
	say()
}

// 打的函数，不管传进来声明，我都要打 Ta，打 Ta 就会叫，就要执行 Ta 的 say 方法
func beat(arg sayer) {
	arg.say()
}

func main() {
	cat1 := cat{}
	beat(cat1)
	dog1 := dog{}
	beat(dog1)
	person1 := person{name: "Bingo"}
	beat(person1)

	var s sayer
	c2 := cat{}
	s = c2
	p2 := person{name: "Gophist"}
	s = p2
	fmt.Println(s)

}

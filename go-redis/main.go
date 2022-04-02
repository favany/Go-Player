package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func main() {
	// 1. 连接到 redis
	// c 是 connect 缩写
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("connect success!", c)
	defer c.Close()

	//c2, err := redis.DialURL(url)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println("connect success!", c2)
	//defer c2.Close()

	// 2. 向 redis 写入数据 string [key - value]
	//_, err = c.Do("Set", "name", "Gophist")
	//if err != nil {
	//	fmt.Println("Set error!,", err)
	//	return
	//}
	//fmt.Println("Set success!")

	// 3. 向 redis 读取数据
	name, err := redis.String(c.Do("Get", "name"))
	if err != nil {
		fmt.Println("Get error!,", err)
		return
	}
	// 因为返回的 r 是 interface()
	// 因为 name 对应的值是 string，因此我们需要转换
	fmt.Println("name Get success! name is", name)
}

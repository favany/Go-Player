package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	for i := 0; i < 10; i++ {
		once.Do(RunOnce)
		fmt.Println("Run RunOnce finished.")
	}
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(goRunOnce)
			fmt.Println("Run goRunOnce finished.")
		}()
	}
}

func RunOnce() {
	fmt.Println("in RunOnce")
}

func goRunOnce() {
	fmt.Println("in goRunOnce")
}

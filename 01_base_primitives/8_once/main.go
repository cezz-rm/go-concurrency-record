package main

import (
	"fmt"
	"sync"
)

// Once 常常用来初始化单利资源
// 或者并发访问只需初始化一次的共享资源
// 或者在测试的时候初始化一次测试资源
func main() {
	var once sync.Once

	f1 := func() {
		fmt.Println("in f1")
	}
	once.Do(f1)

	f2 := func() {
		fmt.Println("in f2")
	}
	once.Do(f2)
}

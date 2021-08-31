package main

import (
	"fmt"
	"sync"
)

// Unlock未加锁的Mutex, 会直接panic
// 1. Lock/Unlock不是成对出现
func foo() {
	var mu sync.Mutex
	defer mu.Unlock()
	fmt.Println("hello")
}

// 2. Copy已使用的Mutex(Mutex是一个有状态的对象)
// 使用vet工具进行检查死锁(实现了Locker接口就会被分析)
// go vet main.go
//type Counter struct {
//	sync.Mutex
//	Count int
//}
//func main() {
//	var c Counter
//	c.Lock()
//	defer c.Unlock()
//	c.Count++
//	foo(c) // 复制锁
//}
//// 这里Counter的参数是通过复制的方式传入的
//func foo(c Counter) {
//	c.Lock()
//	defer c.Unlock()
//	fmt.Println("in foo")
//}

// 3.重入(Mutex不是可重入的锁)

// 4.死锁(互斥、持有和等待、不可剥夺、环路等待)

func main() {
	foo()
}

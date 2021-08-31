package main

import (
	"fmt"
	"sync"
)

// 当一个goroutine调用Lock方法获得了这个锁的拥有权后
// 其他请求锁的goroutine就会阻塞在Lock方法的调用上
// 直到锁被释放并且自己获取到了这个锁的拥有权

// race detector(检测并发访问共享资源是否有问题)
// go run -race main.go
func main() {
	count := 0
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				mu.Lock() // 加锁
				count++
				mu.Unlock() // 释放锁
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

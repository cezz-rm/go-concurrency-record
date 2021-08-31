package main

import (
	"fmt"
	"sync"
	"time"
)

// 使用WaitGroup时的常见错误
var wg sync.WaitGroup

// 1.计数器设置为负值(调用Done方法的次数过多)
func foo1() {
	wg.Add(10)

	wg.Add(-10)

	wg.Add(-1)
}

func foo2() {
	wg.Add(1)

	wg.Done()

	wg.Done()
}

// 2.不期望的Add时机
// 必须等所有的Add方法调用完成之后再调用Wait
func dosomething(millisecs time.Duration, wg *sync.WaitGroup) {
	duration := millisecs * time.Millisecond
	time.Sleep(duration) // 故意sleep一段时间
	wg.Add(1)
	fmt.Println("后台执行, duration:", duration)
	wg.Done()
}

//func main() {
//	go dosomething(100, &wg) // 启动第一个goroutine
//	go dosomething(110, &wg) // 启动第二个goroutine
//	go dosomething(120, &wg) // 启动第三个goroutine
//	go dosomething(130, &wg) // 启动第四个goroutine
//	wg.Wait() // 主goroutine等待完成
//	fmt.Println("Done")
//}

// 3. 前一个Wait还未结束就重用WaitGroup
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done() // 计数器减1
		wg.Add(1) // 计数值加1
	}()
	wg.Wait() // 主goroutine等待，有可能和第7行并发执行
}

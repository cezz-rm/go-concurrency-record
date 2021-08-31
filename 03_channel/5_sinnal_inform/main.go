package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// chan如果为空, receiver接收数据的时候就会阻塞等待
// 直到chan被关闭或者有新的数据到来

func main() {
	closing := make(chan struct{})
	closed := make(chan struct{})

	go func() {
		// 模拟业务处理
		for {
			select {
			case <-closing:
				return
			default:
				// 业务计算
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// 处理ctrl+c中断信号
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	close(closing)
	// 执行退出之前的清理动作
	go doCleanup(closed)

	select {
	case <-closed:
	case <-time.After(time.Second):
		fmt.Println("清理超时, 不等了")
	}
	fmt.Println("优雅退出")
}

func doCleanup(closed chan struct{}) {
	time.Sleep(time.Minute)
	close(closed)
}

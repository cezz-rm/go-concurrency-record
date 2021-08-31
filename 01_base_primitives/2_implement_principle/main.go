package main

import (
	"fmt"
	"sync"
)

// Unlock 方法可以被任意的 goroutine 调用释放锁，
// 即使是没持有这个互斥锁的 goroutine，也可以进行这个操作。

// 这是因为，Mutex 本身并没有包含持有这把锁的 goroutine 的信息，
// 所以，Unlock 也不会对此进行检查。Mutex 的这个设计一直保持至今。

type Foo struct {
	mu    sync.Mutex
	count int
}

func (f *Foo) Bar() {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.count < 1000 {
		f.count += 3
		return
	}

	f.count++
	return
}

func main() {
	f := Foo{count: 1}
	f.Bar()
	fmt.Println(f.count)
}

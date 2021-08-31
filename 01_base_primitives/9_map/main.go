package main

// map的类型是map[key], key类型必须时可比较的
//func main() {
//	m := make(map[string]int)
//	m["a"] = 0
//	fmt.Printf("a=%d, b=%d", m["a"], m["b"])
//}

// 1.必须初始化
//type Counter struct {
//	Website      string
//	Start        time.Time
//	PageCounters map[string]int
//}
//func main() {
//	var c Counter
//	c.Website = "baidu.com"
//	c.PageCounters["/"]++
//}

// 2.并发读写

//func main() {
//	var m = make(map[int]int,10) // 初始化一个map
//	go func() {
//		for {
//			m[1] = 1 //设置key
//		}
//	}()
//	go func() {
//		for {
//			_ = m[2] //访问这个map
//		}
//	}()
//	select {}
//}

// 什么场景使用sync.map, 而不是map+RWMutex
// 1. 只会增长的缓存系统, 一个key只写入一次而被读很多次
// 2. 多个goroutine为不相交的键集读、写和重写键值对

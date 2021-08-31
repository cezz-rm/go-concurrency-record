package main

import (
	"fmt"
	"time"
)

// goroutine只需关注自己信箱(channel), 处理完毕后, 就把结果发送到下一家的信箱中

type Token struct{}

func newWorker(id int, ch chan Token, nextCh chan Token) {
	for {
		token := <-ch
		fmt.Println(id + 1)
		time.Sleep(time.Second)
		nextCh <- token
	}
}

func main() {
	chs := []chan Token{make(chan Token), make(chan Token), make(chan Token), make(chan Token)}

	for i := 0; i < 4; i++ {
		go newWorker(i, chs[i], chs[(i+1)%4])
	}

	chs[0] <- struct{}{}

	select {}
}

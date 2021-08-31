package main

import (
	"context"
	"fmt"
	"time"
)

//https://golangbyexample.com/using-context-in-golang-complete-guide/
func task(ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Gracefully exit")
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second * 1)
			i++
		}
	}
}

func main() {
	ctx := context.Background()
	cancelCrx, cancelFunc := context.WithCancel(ctx)
	go task(cancelCrx)
	time.Sleep(time.Second * 3)
	cancelFunc()
	time.Sleep(time.Second * 1)
}

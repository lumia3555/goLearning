package main

import (
	"context"
	"fmt"
	"time"
)

type otherContext struct {
	context.Context
}

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get msg to cancel\n", name)
			return
		default:
			fmt.Printf("%s is running \n", name)
			time.Sleep(1 * time.Second)
		}
	}
}

func workWithValue(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get msg to cancel\n", name)
			return
		default:
			value := ctx.Value("key").(string)
			fmt.Printf("%s is running value=%s \n", name, value)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// 使用context.Background()构建一个WithCancel类型的上下文
	ctxa, cancel := context.WithCancel(context.Background())

	go work(ctxa, "work a")

	// 使用WithDeadline包装前面的上下文,具有超时通知
	tm := time.Now().Add(3 * time.Second)
	ctxb, _ := context.WithDeadline(ctxa, tm)

	go work(ctxb, "work b")

	// 使用WithValue包装前面的上下文，能够传递数据
	oc := otherContext{ctxb}
	ctxc := context.WithValue(oc, "key", "andes, pass from main")

	go workWithValue(ctxc, "work c")

	time.Sleep(10 * time.Second)

	cancel()

	time.Sleep(5 * time.Second)
	fmt.Println("main stop")
}

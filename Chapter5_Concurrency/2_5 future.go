package main
import (
	"fmt"
	"time"
)

type query struct {
	// 参数channel
	sql chan string

	// 结果channel
	result chan string
}

func execQuery(q query) {
	go func() {
		sql := <-q.sql
		q.result <- "result from " + sql
	}()
}

func main() {
	q := query{
		make(chan string, 1),
		make(chan string, 1),
	}

	go execQuery(q)

	q.sql <- "hello from my phone"

	time.Sleep(3 * time.Second)

	fmt.Println(<-q.result)
}

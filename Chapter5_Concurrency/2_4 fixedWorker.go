package main
import "fmt"

const (
	NUMBER = 10
)

type task struct {
	begin int
	end int
	result chan<-int
}

func (t *task) do() {
	sum := 0
	for i:=t.begin; i<=t.end; i++ {
		sum += i 
	}
	t.result <- sum
}

func InitTask(taskchan chan<-task, r chan int, p int) {
	qu := p/10
	mod := p%10
	high := qu*10

	for j:=0; j<qu; j++ {
		b := 10*j + 1
		e := 10*(j+1)
		tsk := task{
			begin: b,
			end: e,
			result: r,
		}
		taskchan <- tsk
	}

	if mod != 0 {
		tsk := task{
			begin: high + 1,
			end: p,
			result: r,
		}
		taskchan <- tsk
	}
	close(taskchan)
}

// 根据worker的数量生成goroutine
func DistributeTask(taskchan <-chan task, workers int, done chan struct{}) {
	for i:=0; i<workers; i++ {
		go ProcessTask(taskchan, done)
	}
}

func ProcessTask(taskchan <-chan task, done chan struct{}) {
	for t := range taskchan {
		t.do()
	}
	done <- struct{}{}
}

func CloseResult(done chan struct{}, resultchan chan int, workers int) {
	for i:=0; i<workers; i++ {
		<-done
	}
	close(done)
	close(resultchan)
}

func ProcessResult(resultchan chan int) int {
	sum := 0
	for r := range resultchan {
		// fmt.Println(r)
		sum += r 
	}
	return sum
}

func main() {
	workers := NUMBER
	sumTo := 200
	taskchan := make(chan task, workers)
	resultchan := make(chan int, workers)
	done := make(chan struct{}, workers)

	go InitTask(taskchan, resultchan, sumTo)
	go DistributeTask(taskchan, workers, done)
	go CloseResult(done, resultchan, workers)

	sum := ProcessResult(resultchan)
	fmt.Println("sum =", sum)
}
package main
import (
	"fmt"
	"sync"
)

type task struct {
	begin int
	end int
	result chan<-int
}

// task run: calculate sum from begin to end, write result into chan result
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

// taskchan中的每一个task都生成一个goroutine去处理
func DistributeTask(taskchan <-chan task, wait *sync.WaitGroup, result chan int) {
	for v := range taskchan {
		wait.Add(1)
		go ProcessTask(v, wait)
	}
	wait.Wait()
	
	close(result)
}

func ProcessTask(t task, wait *sync.WaitGroup) {
	t.do()
	wait.Done()
}

func ProcessResult(resultchan chan int) int {
	sum := 0
	for r := range resultchan {
		fmt.Println(r)
		sum += r 
	}
	return sum
}

func main() {
	// create task channel
	taskchan := make(chan task, 10)

	// create result channel
	resultchan := make(chan int, 10)

	wait := &sync.WaitGroup{}

	// 构建不同数量的task并分发到各个task通道中，生成taskchan
	go InitTask(taskchan, resultchan, 100)

	// 针对每个taskchan中的task，启动一个goroutine处理任务，将结果放在resultchan中
	go DistributeTask(taskchan, wait, resultchan)

	// 读取resultchan中结果并返回最终值
	sum := ProcessResult(resultchan)

	fmt.Println("sum =", sum)
}

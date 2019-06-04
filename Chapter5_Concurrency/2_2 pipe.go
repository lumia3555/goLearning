package main
import "fmt"

func chain(in chan int) chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- 1 + v
		}
		close(out)
	}()
	return out
}

func main() {
	in := make(chan int)
	go func() {
		for i:=0; i<10; i++ {
			in <- i 
		}
		close(in)
	}()
	// 为什么添加下面打印in里的东西后，out那部分的就不打印了?
	// for v := range in {
	// 	fmt.Printf("%v ", v)
	// }
	// fmt.Printf("\n break \n")
	out := chain(chain(chain(in)))
	for va := range out {
		fmt.Printf("%v-", va)
	}
}
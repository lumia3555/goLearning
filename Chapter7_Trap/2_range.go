package main
import "sync"

func main() {
	wg := sync.WaitGroup{}
	si := []int{1,2,3,4,5,6,7}

	for i := range si {
		wg.Add(1)
		go func(l int) {
			println(l)
			wg.Done()
		}(i)
	}

	// var i *int = 0x00c420094010
	// println(*i)

	wg.Wait()
}
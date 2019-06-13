package common

import "fmt"

func Pslice(slice *[]int) {
	for _,v := range *slice {
		fmt.Printf("%v ", v)
	}
	fmt.Printf("\n")
}

func PsliceS(slice *[]string) {
	for _,v := range *slice {
		fmt.Printf("%v ", v)
	}
	fmt.Printf("\n")
}

func Pmap(m *map[string]string) {
	for k,v := range *m {
		fmt.Printf("k -> %v, v -> %v \n", k, v)
	}
}
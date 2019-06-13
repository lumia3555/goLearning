package main

import (
	"./common"
	"fmt"
)

func sliceTest() {
	slice := []int{1,2,3,4,5}
	newSlice := slice[1:3]

	common.Pslice(&slice)
	common.Pslice(&newSlice)

	newSlice = append(newSlice, 99)
	common.Pslice(&slice)
	common.Pslice(&newSlice)
}

func sliceTest2() {
	source := []string{"A","B","C","D","E"}
	// slice[i:j:k]
	// len = j - i
	// cap = k - i
	common.PsliceS(&source)
	appendSlice(source)
	common.PsliceS(&source)

}

func appendSlice(slice []string) []string {
	result := append(slice[0:2], "FFF")
	common.PsliceS(&result)
	return result
}

func mapTest() {
	// nil映射不能用来存储键值对
	// var colors map[string]string
	colors := map[string]string{}
	colors["Red"] = "#da1337"
	common.Pmap(&colors)
}

func floatTest() {
	var a float32 = 0.1
	var b float64 = 0.2
	if a+b == 0.3 {
		fmt.Printf("%v \n", a+b)
	} else {
		fmt.Printf("rr \n")

	}
}


func main() {
	floatTest()
}
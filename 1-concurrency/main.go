package main

import (
	"fmt"
	"math/rand"
)

func main() {
	countNum := 10
	ch := make(chan int)
	go createSliceOfRandomElements(countNum, ch)
	for i := 0; i < countNum; i++ {
		fmt.Printf("%d ", <-ch)
	}
}

func createSliceOfRandomElements(count int, ch chan int) {
	sliceOfRandomElements := make([]int, 10)
	for i := 0; i < count; i++ {
		sliceOfRandomElements[i] = rand.Intn(101)
	}
	for _, val := range sliceOfRandomElements {
		go squareNumber(val, ch)
	}
}

func squareNumber(val int, ch chan int) {
	res := val * val
	ch <- res
}

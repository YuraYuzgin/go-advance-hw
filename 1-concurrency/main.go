package main

import (
	"fmt"
	"math/rand"
)

func main() {
	countNum := 10
	ch1 := make(chan int)
	ch2 := make(chan int)
	go createSliceOfRandomElements(countNum, ch1)
	go squareNumber(ch1, ch2)
	for range countNum {
		fmt.Printf("%d ", <-ch2)
	}
}

func createSliceOfRandomElements(count int, ch1 chan int) {
	defer close(ch1)
	sliceOfRandomElements := make([]int, count)
	for i := range count {
		sliceOfRandomElements[i] = rand.Intn(101)
	}
	for _, val := range sliceOfRandomElements {
		ch1 <- val
	}
}

func squareNumber(ch1 chan int, ch2 chan int) {
	defer close(ch2)
	for val := range ch1 {
		ch2 <- val * val
	}
}

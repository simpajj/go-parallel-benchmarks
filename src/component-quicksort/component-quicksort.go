package main

import (
	_ "fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"
)

const RANGE = 1000

func qsort_pass(arr []int, done chan int) []int {
	if len(arr) < 2 {
		done <- len(arr)
		return arr
	}
	pivot := arr[0]
	i, j := 1, len(arr)-1
	for i != j {
		for arr[i] < pivot && i != j {
			i++
		}
		for arr[j] >= pivot && i != j {
			j--
		}
		if arr[i] > arr[j] {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	if arr[j] >= pivot {
		j--
	}
	arr[0], arr[j] = arr[j], arr[0]
	done <- 1

	go qsort_pass(arr[:j], done)
	go qsort_pass(arr[j+1:], done)
	return arr
}

func qsort(arr []int) []int {
	done := make(chan int)
	// defer func() {
	// 	close(done)
	// }()

	go qsort_pass(arr[:], done)

	result := len(arr)
	for result > 0 {
		result -= <-done
	}
	return arr
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	N, _ := strconv.Atoi(os.Args[1])
	iCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(iCPU)

	arr := make([]int, RANGE)
	for i := 0; i <= N; i++ {
		for j := range arr {
			arr[j] = rand.Intn(RANGE)
		}
		qsort(arr)
	}
}

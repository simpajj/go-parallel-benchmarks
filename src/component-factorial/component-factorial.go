package main

import (
	_ "fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"
)

const (
	RANGE int = 1000 // The number of elements in the array
)

/* Recursive factorial function */
func factorial(n uint64) uint64 {
	if n > 1 {
		n = n * factorial(n-1)
	} else {
		return uint64(1)
	}
	return n
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	N, _ := strconv.Atoi(os.Args[1]) // Iterations
	iCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(iCPU)

	arr := make([]int, RANGE)
	for i := 0; i < RANGE; i++ {
		arr[i] = rand.Intn(10 - 0)
	}

	for i := 0; i <= N; i++ {
		for j := 0; j < RANGE; j++ {
			// go fmt.Printf("Factorial of: %d = %d\n", arr[j], factorial(uint64(arr[j])))
			go factorial(uint64(arr[j]))
		}
	}
}

package main

import (
	_ "fmt"
	"os"
	"runtime"
	"strconv"
)

const (
	RANGE int = 1000
)

func factorial(n uint64) uint64 {
	if n > 1 {
		n = n * factorial(n-1)
	} else {
		return uint64(1)
	}
	return n
}

func main() {
	N, _ := strconv.Atoi(os.Args[1]) // Iterations
	iCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(iCPU)

	for i := 0; i <= N; i++ {
		for j := 0; j < RANGE; j++ {
			go factorial(uint64(j))
		}
	}
}

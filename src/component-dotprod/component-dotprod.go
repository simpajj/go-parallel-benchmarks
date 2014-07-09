package main

import (
	_ "fmt"
	"os"
	"runtime"
	"strconv"
)

func dotprod(A, B []float64, result float64) float64 {
	for i := range A {
		result += A[i] * B[i]
	}
	return result
}

func main() {
	N, _ := strconv.Atoi(os.Args[1]) // Iterations
	SIZE, _ := strconv.Atoi(os.Args[2])
	iCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(iCPU)
	A, B := make([]float64, SIZE), make([]float64, SIZE)

	result := 0.0
	for i := 0; i <= N; i++ {
		for j := range A {
			A[j] = float64(j) * 1.0
			B[j] = float64(j) * 2.0
			go dotprod(A, B, result)
		}
	}
}

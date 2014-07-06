package main

import (
	_ "fmt"
	"os"
	"strconv"
)

const n = 100

func dotprod(A, B []float64, result float64) float64 {
	for i := 0; i < n; i++ {
		result += A[i] * B[i]
	}
	// fmt.Printf("%f\n", result)
	return result
}

func main() {
	N, _ := strconv.Atoi(os.Args[1]) // Iterations
	A, B := make([]float64, 100), make([]float64, 100)

	result := 0.0
	for i := 0; i <= N; i++ {
		for j := 0; j < n; j++ {
			A[j] = float64(j) * 1.0
			B[j] = float64(j) * 2.0
			go dotprod(A, B, result)
		}
	}
	// fmt.Printf("%f", result)
}

package main

import (
	_ "fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const RANGE = 5
const (
	ROWS = 3
	COLS = 3
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var sum int = 0
	N, _ := strconv.Atoi(os.Args[1])

	/* Create matrices */
	first := make([][]int, ROWS)
	second := make([][]int, ROWS)
	result := make([][]int, ROWS)
	for i := range first {
		first[i] = make([]int, COLS)
		second[i] = make([]int, COLS)
		result[i] = make([]int, COLS)
	}

	for i := 0; i <= N; i++ {
		/* Populate matrices */
		for c := 0; c < ROWS; c++ {
			for d := 0; d < COLS; d++ {
				first[c][d] = rand.Intn(RANGE - 0)
			}
		}

		for c := 0; c < ROWS; c++ {
			for d := 0; d < COLS; d++ {
				second[c][d] = rand.Intn(RANGE - 0)
			}
		}

		/* Multiply matrices */
		for c := 0; c < ROWS; c++ {
			for d := 0; d < COLS; d++ {
				for k := 0; k < ROWS; k++ {
					sum = sum + first[c][k]*second[k][d]
				}

				result[c][d] = sum
				sum = 0
			}
		}

		// fmt.Println("Product: \n")
		// for c := 0; c < ROWS; c++ {
		// 	for d := 0; d < COLS; d++ {
		// 		fmt.Printf("%d\t", result[c][d])
		// 	}
		// 	fmt.Printf("\n")
		// }
	}
}

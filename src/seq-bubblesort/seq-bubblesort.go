package main

import (
	_ "fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	N, _ := strconv.Atoi(os.Args[1])
	number_of_elements, _ := strconv.Atoi(os.Args[2])
	array := make([]int, number_of_elements)
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i <= N; i++ {
		for c := 0; c < number_of_elements; c++ {
			array[c] = rand.Intn(100 - 0)
		}

		bubble_sort(array, number_of_elements)
	}
}

func bubble_sort(list []int, n int) {
	for c := 0; c < (n - 1); c++ {
		for d := 0; d < (n - 1); d++ {
			if list[d] > list[d+1] {
				t := list[d]
				list[d] = list[d+1]
				list[d+1] = t
			}
		}
	}
}

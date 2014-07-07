package main

import (
	_ "fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const number_of_elements = 10000000

func main() {
	array := make([]int, number_of_elements)
	rand.Seed(time.Now().UTC().UnixNano())
	N, _ := strconv.Atoi(os.Args[1])

	for i := 0; i <= N; i++ {
		for c := 0; c < number_of_elements; c++ {
			array[c] = rand.Intn(100 - 0)
		}

		bubble_sort(array, number_of_elements)
		// fmt.Printf("Sorted list:\n")

		// for c := 0; c < number_of_elements; c++ {
		// 	fmt.Printf("%d\n", array[c])
		// }
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

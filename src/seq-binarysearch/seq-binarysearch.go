package main

import (
	_ "fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const number_of_elements = 100

func binarysearch(array []int) int {
	search := 13
	for c := range array {
		array[c] = rand.Intn(number_of_elements - 0)
	}

	first := 0
	last := number_of_elements - 1
	middle := (first + last) / 2

	for first <= last {
		if array[middle] < search {
			first = middle + 1
		} else if array[middle] == search {
			search = array[middle]
			break
		} else {
			last = middle - 1
		}

		middle = (first + last) / 2
	}

	if first > last {
		return -1
	}
	return search
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	N, _ := strconv.Atoi(os.Args[1])
	array := make([]int, number_of_elements)

	for i := 0; i <= N; i++ {
		binarysearch(array)
	}
}

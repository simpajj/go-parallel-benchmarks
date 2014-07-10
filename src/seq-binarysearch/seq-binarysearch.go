package main

import (
	_ "fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const number_of_elements = 100

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	N, _ := strconv.Atoi(os.Args[1])
	var first, last, middle int
	search := 13
	array := make([]int, number_of_elements)

	for i := 0; i <= N; i++ {
		for c := range array {
			array[c] = rand.Intn(number_of_elements - 0)
		}

		first = 0
		last = number_of_elements - 1
		middle = (first + last) / 2

		for first <= last {
			if array[middle] < search {
				first = middle + 1
			} else if array[middle] == search {
				break
			} else {
				last = middle - 1
			}

			middle = (first + last) / 2
		}

		if first > last {
			break
		}
	}
}

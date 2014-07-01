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
	var first, last, middle int
	search := 13
	array := make([]int, 100)
	N, _ := strconv.Atoi(os.Args[1])

	for i := 0; i <= N; i++ {
		for c := 0; c < number_of_elements; c++ {
			array[c] = rand.Intn(100 - 0)
		}

		first = 0
		last = number_of_elements - 1
		middle = (first + last) / 2

		for first <= last {
			if array[middle] < search {
				first = middle + 1
			} else if array[middle] == search {
				// fmt.Printf("%d found at location %d \n", search, middle+1)
				os.Exit(0)
			} else {
				last = middle - 1
			}

			middle = (first + last) / 2
		}

		if first > last {
			// fmt.Println("The number is not in the list!")
			os.Exit(0)
		}
	}
}

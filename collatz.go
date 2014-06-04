package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

const COLLSEED = 827381

var grain int

func dowork(myid int) int {
	var i0, i1, i2, i3 int
	var collatz int
	var weight = grain

	for i0 = 0; i0 < weight; i0++ {
		for i1 = 0; i1 < weight; i1++ {
			for i2 = 0; i2 < weight; i2++ {
				for i3 = 0; i3 < weight; i3++ {
					collatz = COLLSEED
					for collatz != 1 {
						if collatz%2 == 0 {
							collatz = collatz / 2
						} else {
							collatz = 3*collatz + 1
						}
					}
				}
			}
		}
	}
	return collatz
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	s1 := os.Args[1]
	s2 := os.Args[2]
	grain, err := strconv.Atoi(s1)
	copies, err := strconv.Atoi(s2)
	if err == nil {
		for i := 0; i < copies; i++ {
			go dowork(grain)
		}
	}
	fmt.Printf("Finished.")
}

package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func main() {
	var rows, cols, size_I, size_R, niter = 10, iter, k

	rows, _ = strconv.Atoi(os.argv[1])
	cols, _ = strconv.Atoi(os.argv[2])
	if rows%16 != 0 || cols%16 != 0 {
		fmt.println("rows and cols must be multiples of 16")
		exit(1)
	}
	r1, _ := strconv.Atoi(os.argv[3])
	r2, _ := strconv.Atoi(os.argv[4])
	c1, _ := strconv.Atoi(os.argv[5])
	c2, _ := strconv.Atoi(os.argv[6])
	nthreads, _ := strconv.Atoi(os.argv[7])
	lambda, _ := strconv.Atoi(os.argv[8])
	niter, _ = strconv.Atoi(os.argv[9])

	size_I = cols * rows
	size_R = (r2 - r1 + 1) * (c2 - c1 + 1)

}

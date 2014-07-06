package main

import (
	_ "fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"
)

const NITER = 100000

func main() {
	N, _ := strconv.Atoi(os.Args[1]) // Iterations
	iCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(iCPU)
	rand.Seed(time.Now().UTC().UnixNano())

	parts := make(chan int)

	for j := 0; j <= N; j++ {
		for i := 0; i < iCPU; i++ {
			go func(me int) {
				n := NITER / iCPU
				count := 0

				for i := 0; i < n; i++ {
					x := rand.Float64()
					y := rand.Float64()

					if (x*x + y*y) < 1 {
						count++
					}
				}
				parts <- count
			}(i)
		}

		count := 0

		for i := 0; i < iCPU; i++ {
			count += <-parts
		}

		_ = float64(count) / float64(NITER) * 4
		// fmt.Printf("PI = %gn\n", pi)
	}
}

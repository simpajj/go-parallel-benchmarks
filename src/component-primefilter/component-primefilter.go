package main

import (
	_ "fmt"
	"os"
	"runtime"
	"strconv"
)

const RANGE = 50000

func generate(ch chan int) {
	go func() {
		for i := 2; i <= RANGE; i++ {
			ch <- i
		}
		close(ch)
	}()
}

func filter(in chan int, out chan int, filter int) {
	for val := range in {
		if val%filter != 0 {
			out <- val
		}
	}
	close(out)
}

func main() {
	N, _ := strconv.Atoi(os.Args[1])
	iCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(iCPU)

	for i := 0; i <= N; i++ {
		ch := make(chan int, 1)
		generate(ch)
		for _ = range ch {
			prime := <-ch
			ch1 := make(chan int, 1)
			go filter(ch, ch1, prime)
			ch = ch1
		}
	}
}

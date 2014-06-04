package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func fib(n int) int {
	if n < 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func printFib(n int, ch chan<- bool) {
	fmt.Printf("fib %d is %d\n", n, fib(n))
	ch <- true
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	s1 := os.Args[1]
	s2 := os.Args[2]
	n, err := strconv.Atoi(s1)
	copies, err := strconv.Atoi(s2)

	ch := make(chan bool)
	if err == nil {
		for i := 0; i < copies; i++ {
			go printFib(n, ch)
		}
	}

	for i := 0; i < copies; i++ {
		<-ch

	}
}

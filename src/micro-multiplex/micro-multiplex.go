package main

import (
	"os"
	"runtime"
	"strconv"
	"sync"
)

/* Multiplexes multiple values onto a single channel */
func multiplex(inputs []<-chan int, output chan<- int) {
	var group sync.WaitGroup
	for i := range inputs {
		group.Add(1)
		go func(input <-chan int) {
			for val := range input {
				output <- val
			}
			group.Done()
		}(inputs[i])
	}
	go func() {
		group.Wait() // Wait for the last input until we close the output channel
		close(output)
	}()
}

func main() {
	copies, err := strconv.Atoi(os.Args[1])
	cores, err := strconv.Atoi(os.Args[2])
	N, err := strconv.Atoi(os.Args[3])
	runtime.GOMAXPROCS(cores)

	inputs := make([]<-chan int, copies)
	output := make(chan int)
	for i := 0; i <= N; i++ {
		if err == nil {
			go multiplex(inputs, output)
		}
	}
}

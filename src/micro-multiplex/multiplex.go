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
	s1 := os.Args[1]
	s2 := os.Args[2]
	copies, err := strconv.Atoi(s1)
	n, err := strconv.Atoi(s2)
	runtime.GOMAXPROCS(n)

	inputs := make([]<-chan int, copies)
	output := make(chan int)
	if err == nil {
		multiplex(inputs, output)
	}
}

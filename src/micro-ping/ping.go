package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func main() {
	s1 := os.Args[1] // Number of goroutines
	s2 := os.Args[2] // Number of CPUs
	copies, err := strconv.Atoi(s1)
	n, err := strconv.Atoi(s2)
	runtime.GOMAXPROCS(n)

	messages := make(chan string)

	if err == nil {
		for i := 0; i < copies; i++ {
			go func() { messages <- "ping" }()

			msg := <-messages
			fmt.Println(msg)
		}
	}
}

package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func main() {
	copies, err := strconv.Atoi(os.Args[1])
	cores, err := strconv.Atoi(os.Args[2])
	N, err := strconv.Atoi(os.Args[3])
	runtime.GOMAXPROCS(cores)

	messages := make(chan string)

	for i := 0; i <= N; i++ {
		if err == nil {
			for i := 0; i < copies; i++ {
				go func() { messages <- "ping" }()

				msg := <-messages
				fmt.Println(msg)
			}
		}
	}
}

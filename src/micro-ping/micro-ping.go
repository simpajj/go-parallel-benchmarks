package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func main() {
	N, err := strconv.Atoi(os.Args[1]) // Iterations
	iCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(iCPU)

	messages := make(chan string)

	for i := 0; i <= N; i++ {
		if err == nil {
			for i := 0; i < iCPU; i++ {
				go func() { messages <- "ping" }()

				msg := <-messages
				fmt.Println(msg)
			}
		}
	}
}

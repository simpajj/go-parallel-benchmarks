package main

import (
	_ "fmt"
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
			go func() { messages <- "ping" }()
			go func() {
				msg := <-messages
				_ = msg
				// fmt.Println(msg)
			}()
		}
	}
}

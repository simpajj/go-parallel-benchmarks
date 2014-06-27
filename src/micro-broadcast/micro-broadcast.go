package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func broadcast(c chan<- string, n int) {
	s := "broadcasting"
	for i := 0; i <= n; i++ {
		c <- s
	}
}

func listen(c <-chan string) {
	msg := <-c
	fmt.Println(msg)
}

func main() {
	goroutines, err := strconv.Atoi(os.Args[1])
	cores, err := strconv.Atoi(os.Args[2])
	N, err := strconv.Atoi(os.Args[3])
	runtime.GOMAXPROCS(cores)

	broadcast_chan := make(chan string, goroutines)

	for i := 0; i <= N; i++ {
		go broadcast(broadcast_chan, goroutines)

		if err == nil {
			for i := 0; i <= goroutines; i++ {
				go listen(broadcast_chan)
			}
		}
	}
}

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
	s1 := os.Args[1]
	s2 := os.Args[2]
	goroutines, err := strconv.Atoi(s1)
	cores, err := strconv.Atoi(s2)
	runtime.GOMAXPROCS(cores)

	broadcast_chan := make(chan string, goroutines)
	// Remove go so it's not parallel?
	go broadcast(broadcast_chan, goroutines)

	if err == nil {
		for i := 0; i <= goroutines; i++ {
			go listen(broadcast_chan)
		}
	}
}

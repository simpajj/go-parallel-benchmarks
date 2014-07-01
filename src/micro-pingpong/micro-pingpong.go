package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

/* Sends pings */
func ping(ping chan<- string, msg string) {
	ping <- msg
}

/* Receives pings and sends pongs */
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	msg = "pong"
	pongs <- msg
}

func main() {
	copies := 2
	N, err := strconv.Atoi(os.Args[1])
	runtime.GOMAXPROCS(copies)

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	for i := 0; i <= N; i++ {
		if err == nil {
			for i := 0; i < copies; i++ {
				go ping(pings, "ping")
				go pong(pings, pongs)
				fmt.Println(<-pongs)
			}
		}
	}
}

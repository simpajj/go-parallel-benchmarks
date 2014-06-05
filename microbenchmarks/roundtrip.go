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
	s1 := os.Args[1]
	s2 := os.Args[2]
	copies, err := strconv.Atoi(s1)
	n, err := strconv.Atoi(s2)
	runtime.GOMAXPROCS(n)

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	if err == nil {
		for i := 0; i < copies; i++ {
			go ping(pings, "ping")
			go pong(pings, pongs)
			fmt.Println(<-pongs)
		}
	}
}

package main

import (
	_ "fmt"
	"os"
	"runtime"
	"strconv"
)

/* Sends pings, should do the same as pong in other order */
func pingpong(ping chan<- string, msg string, pongs <-chan string) {
	ping <- msg
	_ = <-pongs
	// fmt.Println(<-pongs)
}

/* Receives pings and sends pongs */
func pongping(pings <-chan string, pongs chan<- string) {
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
			go pingpong(pings, "ping", pongs)
			go pongping(pings, pongs)
		}
	}
}

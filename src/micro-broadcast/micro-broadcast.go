package main

import (
	_ "fmt"
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
	_ = msg
	// fmt.Println(msg)
}

func main() {
	N, err := strconv.Atoi(os.Args[1]) // Iterations
	iCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(iCPU)

	broadcast_chan := make(chan string, iCPU)

	for i := 0; i <= N; i++ {
		go broadcast(broadcast_chan, iCPU)

		if err == nil {
			for i := 0; i <= iCPU; i++ {
				go listen(broadcast_chan)
			}
		}
	}
}

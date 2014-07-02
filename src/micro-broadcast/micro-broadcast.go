package main

import (
	_ "fmt"
	"os"
	"runtime"
	"strconv"
)

func broadcast(c chan<- string, n int) {
	s := "broadcasted"
	for i := 0; i <= n; i++ {
		c <- s
	}
}

func listen(c <-chan string) {
	select {
	case msg, ok := <-c:
		if ok {
			_ = msg
			// fmt.Printf("%s\n", msg)
		} else {
			// fmt.Println("Channel closed!")
			break
		}
	default:
		// fmt.Println("No value ready.")
		break
	}
}

func main() {
	N, err := strconv.Atoi(os.Args[1]) // Iterations
	iCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(iCPU)

	broadcast_chan := make(chan string, iCPU)

	for i := 0; i <= N; i++ {
		if err == nil {
			go broadcast(broadcast_chan, iCPU)
			go listen(broadcast_chan)
		}
	}
}

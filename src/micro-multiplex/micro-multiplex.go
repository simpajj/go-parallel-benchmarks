package main

import (
	_ "fmt"
	"os"
	"runtime"
	"strconv"
)

func broadcast(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			// c <- fmt.Sprintf("%s %d", msg, i)
			continue
		}
	}()
	return c
}

func multiplex(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func main() {
	N, err := strconv.Atoi(os.Args[1]) // Iterations
	iCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(iCPU)

	for i := 0; i <= N; i++ {
		if err == nil {
			_ = multiplex(broadcast("hello"), broadcast("world"))
			// fmt.Println(<-c)
		}
	}
}

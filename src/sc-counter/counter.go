package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func main() {
	s1 := os.Args[1]
	s2 := os.Args[2]
	s3 := os.Args[3]
	copies, err := strconv.Atoi(s1)
	n, err := strconv.Atoi(s2)
	count, err := strconv.Atoi(s3)

	runtime.GOMAXPROCS(n)

	channel := make(chan int)
	for i := 0; i < copies; i++ {
		go func(i int) {
			for v := range channel {
				fmt.Printf("Count %d from goroutine %d\n", v, i)
			}
		}(i)
	}
	if err == nil {
		for i := 0; i <= count; i++ {
			channel <- i
		}
	}

	close(channel)

}

package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

func p() {
	for number := 1; number < 1000000; number++ {
		fmt.Printf("%d ", number)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	fmt.Println("Start")
	s1 := os.Args[1]
	copies, err := strconv.Atoi(s1)
	if err == nil {
		for i := 0; i < copies; i++ {
			go p()
		}
	}
	time.Sleep(time.Second * 1)
	fmt.Println("Finish")
}

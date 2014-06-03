package main

import (
	"fmt"
	"runtime"
	"time"
)

func p() {
	for number := 1; number < 1000000; number++ {
		fmt.Printf("%d ", number)
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	fmt.Println("Start")
	go p()
	go p()
	time.Sleep(1 * time.Second)
	fmt.Println("Finish")
}

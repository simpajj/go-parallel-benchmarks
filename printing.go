package main

import (
	"fmt"
	"time"
)

func p() {
	for number := 1; number < 1000000; number++ {
		fmt.Printf("%d ", number)
	}
}

func main() {
	fmt.Println("Start")
	go p()
	go p()
	time.Sleep(1 * time.Second)
	fmt.Println("Finish")
}

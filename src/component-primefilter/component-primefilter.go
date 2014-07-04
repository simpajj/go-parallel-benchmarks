package main

import _ "fmt"

const RANGE = 50000

func generate(ch chan int) {
	go func() {
		for i := 2; i <= RANGE; i++ {
			ch <- i
		}
		close(ch)
	}()
}

func filter(in chan int, out chan int, filter int) {
	for val := range in {
		if val%filter != 0 {
			out <- val
		}
	}
	close(out)
}

func main() {
	ch := make(chan int, 1)
	generate(ch)

	for _ = range ch {
		prime := <-ch
		// fmt.Print(prime, "\n")
		ch1 := make(chan int, 1)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const RANGE = 1000000

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var first, second, hcf, lcm int

	first = rand.Intn(RANGE - 0)
	second = rand.Intn(RANGE - 0)

	hcf = gcd(first, second)
	lcm = (first * second) / hcf

	fmt.Printf("The greatest common divisor of %d and %d = %d \n", first, second, hcf)
	fmt.Printf("The least common multiple of %d and %d = %d \n", first, second, lcm)
}

func gcd(first, second int) int {
	if second == 0 {
		return first
	} else {
		return gcd(second, first%second)
	}
}

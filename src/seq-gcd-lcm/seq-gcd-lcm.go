package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const RANGE = 1000

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var first, second, hcf, lcm int
	N, _ := strconv.Atoi(os.Args[1])

	for i := 0; i <= N; i++ {
		first = rand.Intn(RANGE - 0)
		second = rand.Intn(RANGE - 0)

		hcf = gcd(first, second)
		lcm = (first * second) / hcf

		fmt.Printf("The greatest common divisor of %d and %d = %d \n", first, second, hcf)
		fmt.Printf("The least common multiple of %d and %d = %d \n", first, second, lcm)
	}
}

func gcd(first, second int) int {
	if second == 0 {
		return first
	} else {
		return gcd(second, first%second)
	}
}

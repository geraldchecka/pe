package main

import (
	"fmt"
	"math"
)

func Algorithm1() int {
	index, isPrime := 1, true
	for i := 3; i <= i; i += 2 {
		isPrime = true
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isPrime = false
				continue
			}
		}
		if isPrime {
			index++
			if index == 10001 {
				return i
			}
		}

	}
	return 0
}
func main() {
	fmt.Printf("Hello Problem 7: %v", Algorithm1())
}

package main

import "fmt"

func Problem1V2() int {
	var sum int = 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	fmt.Printf("Result: %v\n", Problem1V2())
}

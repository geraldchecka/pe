package main

import (
	"fmt"
)

func problem1() int {
	sum1, sum2, sum3 := 0, 0, 0
	var i int = 1
	for (5 * i) < 1000 {
		sum1 += (5*i)
		i++
	}
	i = 1
	for (3 * i) < 1000 {
		sum2 += (3*i)
		i++
	}
	i = 1
	for (15 * i) < 1000 {
		sum3 += (15 * i)
		i++
	}
	return (sum1+sum2-sum3)
}

func main() {
	fmt.Println("Sum of multiples of 3 or 5 less than 1000 is: ",problem1())
}

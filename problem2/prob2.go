package main

import "fmt"
/*I did research on the Internet for a mathematical solution to avoid % comparison */
func FibonacciAlgorithm1() int32 {
	var sum, a, b int32 = 0, 1, 1
	for ; sum < 4000000; {
		sum += (a + b)
		a, b = (a + 2*b), (2*a + 3*b)
	}
	return sum
}
/*This is the original one I wrote*/
func FibonacciAlgorithm2() int32 {
	var a, b, sum, temp int32 = 1, 2, 2, 0 
	for ; sum < 4000000; {
		temp = a + b
		a = b
		b = temp 
		if temp%2 == 0 {
			sum += temp
		}
		if temp >= 4000000 {
			break
		}
	}
	return sum
}

func main() {
	fmt.Printf("Fibonacci result A1: %v\n", FibonacciAlgorithm1())
	fmt.Printf("Fibonacci result A2: %v\n", FibonacciAlgorithm2())
}

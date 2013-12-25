package main

import (
	"fmt"
	"math"
)
//Algorithm 1
func Algorithm1(x uint64) uint64 {
	var lpf uint64
	var i, val uint64 = 2, uint64(math.Sqrt(float64(x)))
	for ; i <= val; i++ {
		if IsPrime(i) {
			if x % i == 0 {
				lpf = i
			} 		
		}
	}
	if lpf == 0 {
		return x
	} else {
		return lpf
	}
}

func IsPrime(x uint64) bool {
	var status = true;
	var i, cmp uint64 = 2, uint64(math.Sqrt(float64(x)))
	for ; i <= cmp; i++ {
		if x%i == 0 {
		    status = false;
		}	
	}
	return status;
}

//Algorithm 2
//Initially I tried this using recursive functions, but that didn't work
func Algorithm2(x int) int {
	var lpf int = 0
	for i := 2; i <= x; i++ {
	    if IsPrime(uint64(i)) {
		for x % i == 0 {
		    x = x / i
		}
	    }
	    lpf = i
	}
	return lpf
}

func main() {
	fmt.Printf("Largest Prime number: A1 - %v\n", Algorithm1(600851475143));
	fmt.Printf("Largest Prime number: A2 - %v\n", Algorithm2(600851475143));
}

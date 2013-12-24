package main

import (
	"fmt"
	"math"
)
/*
	Alogrithm 1
This problem with this algorithm is that, it takes forever to compute large numbers.
*/

//Function to loop through number
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

//Function to validate a number
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

/*
	Algorithm 2
*/
func Algorithm2(x int) int {
	var prime, i int = 0, 2 
	for i <= x/2 {
		fmt.Println(i)
		if IsPrime(uint64(i)) && (x % i == 0) {
			prime = x / i;
			fmt.Println(prime)
			Algorithm2(prime);
		} else {
			i += 1
		}
	}
	return prime
}

func main() {
	fmt.Printf("Largest Prime number:%v\n",Algorithm1(600851475143));
	//fmt.Printf("largest Prime number: %v\n", Algorithm2(15));
}

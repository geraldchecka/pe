package main

import (
	"fmt"
	"math"
)

func IsPrime(x float64) bool {
	var status = true;
	for i, cmp := 2.0, math.Sqrt(x); i <= cmp; i++ {
		if x/i == 0 {
		    status = false;
		}	
	}
	return status;
}

func main() {
	fmt.Printf("Is a prime number:%v\n",IsPrime(10));
}

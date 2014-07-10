package main

import (
	"fmt"
	"math"
)

//Use struct in this alg
func Algorithm2() {

}

func Algorithm1(n int) int {
	return int(math.Pow(float64(n*(n+1)/2), 2) - float64(n*(n+1)*(2*n+1))/6)
}

func main() {
	fmt.Printf("%v\n", Algorithm1(100))
}

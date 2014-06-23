package main

import (
	"fmt"
)
func Algorithm1() int {
	upperNum := 20
	for i := 20; i <= upperNum; i += 20 {
		isMyNumber := true
		for j := 20; j >= 2; j-- {
			if upperNum % j != 0 {
				isMyNumber = false
				break
			}
		}
		if isMyNumber {
			return upperNum
		}
		upperNum += 20
	}
	return upperNum
}

func Algorithm2() string {
	return "Anvesh 2"
}

func main() {
	fmt.Printf("Value is %v\n",Algorithm1())
}

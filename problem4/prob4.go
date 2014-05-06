package main

import (
	"fmt"
	"strconv"
)
func LargestPalindrome() string {
	for	i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			IsPalindrome(strconv.Itoa(i * j));
		}
	}
	return "Anvesh"
}

func IsPalindrome(str string) bool {
	makeString := make([]string, len(str))
	return str == reverse(str)
}

func Reverse(str string) string {
	var revStr string;
	for 
	return revStr
}

func main() {
	fmt.Println("Project Euler - 4\n")
	fmt.Printf("The largest palindrome: %v\n",LargestPalindrome())
}

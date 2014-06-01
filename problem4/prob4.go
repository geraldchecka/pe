package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
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
	return str == Reverse(str)
}

func Reverse(str string) string {
	revStrRune, rIndex := make([]rune, len(str), utf8.RuneCountInString(str)), 0
	for _, rValue := range str {
		revStrRune[rIndex] = rValue
		rIndex++
	}
	for i, rCount := 0, utf8.RuneCountInString(str); i < rCount / 2; i++ {
		revStrRune[i], revStrRune[rCount - 1 - i] = revStrRune[rCount - 1 - i], revStrRune[i]
	}
	return string(revStrRune)
}

func main() {
	fmt.Println("Project Euler - 4")
	//fmt.Printf("The largest palindrome: %v\n",LargestPalindrome())
	fmt.Println(IsPalindrome("sahas"))
}

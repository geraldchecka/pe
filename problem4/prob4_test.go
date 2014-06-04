package main

import (
	"testing"
)

func TestAlgorithm1 (t *testing.T) {
	result1 := LargestPalindrome();
	if result1 != 906609 {
		t.Errorf("Algorithm1: Expected result: %v, Obtained result: %v", 906609, result1)
	}
}

func BenchmarkAlgorithm1 (b *testing.B) {
	for i := 0; i < b.N; i++ {
		LargestPalindrome();
	}
}

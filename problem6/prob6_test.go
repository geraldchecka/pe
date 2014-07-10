package main

import (
	"testing"
)

func TestAlgorithm1(t *testing.T) {
	result1 := Algorithm1(100)
	if result1 != 25164150 {
		t.Errorf("Algorithm1: Expected result: %v, Obtained result: %v", 25164150, result1)
	}
}

func BenchmarkAlgorithm1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Algorithm1(100)
	}
}

package main

import (
	"testing"
)

func TestAlgorithm1(t *testing.T) {
	result1 = Algorithm1()
	if result1 != 232792560 {
		t.Errorf("Algorithm1: Expected result: %v, Obtained result: %v", 232792560, result1)
	}
}

func BenchmarkAlgorithm1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Algorithm1()
	}
}

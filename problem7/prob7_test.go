package main

import (
	"testing"
)

func TestAlgorithm1(t *testing.T) {
	result := Algorithm1()
	if result != 104743 {
		t.Errorf("Algorithm1: Expected result: %v, Obtained result: %v",104743, Algorithm1())
	}
}

func BenchmarkAlgorithm1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Algorithm1()
	}
}

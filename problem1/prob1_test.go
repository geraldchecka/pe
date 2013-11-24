package main

import (
	"testing"
)

func TestProblem1(t *testing.T) {
	result := Problem1()
	if result != 233168 {
		t.Errorf("Algorithm 1: Expected answer: %v, obtained result: %v\n", 233168, result)
	}
}

func TestProblem1V2(t *testing.T) {
	result2 := Problem1V2()
	if result2 != 233168 {
		t.Errorf("ALgorithm 2: Expected answer: %v, obtained result: %v\n", 233168, result2)
	}
}

func BenchmarkAlgorithm1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem1()
	}
}

func BenchmarkAlgorithm2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem1V2()
	}
}

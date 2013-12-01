package main

import (
	"testing"
)

func TestProblem2V1(t *testing.T) {
	result1 := FibonacciAlgorithm1()
	if result1 != 4613732 {
		t.Errorf("Algorithm 1: Expected answer: %v, Obtained result: %v\n",4613732, result1)
	}
}

func TestProblem2V2(t *testing.T) {
	result2 := FibonacciAlgorithm2()
	if result2 != 4613732 {
		t.Errorf("Algorithm2: Expected answer: %v, Obtained result: %v\n", 4613732, result2)
	}
}
func BenchmarkAlgorithm1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciAlgorithm1()
	}
}
func BenchmarkAlgorithm2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciAlgorithm2()
	}
}

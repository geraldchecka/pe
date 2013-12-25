package main

import (
	"testing"
)

func TestAlgorithm1(t *testing.T) {
    result1 := Algorithm1(600851475143)
    if result1 != 6857 {
	t.Errorf("Algorithm1: Expected result: %v, Obtained result: %v",6857, result1)
    }
}

func TestAlgorithm2(t *testing.T) {
    result2 := Algorithm2(600851475143)
    if result2 != 6857 {
	t.Errorf("Algorithm1: Expected result: %v, Obtained result: %v",6857, result2)
    }
}

func BenchmarkAlgorithm1(b *testing.B) {
    for i := 0; i < b.N; i++ {
	Algorithm1(600851475143)
    }
}

func BenchmarkAlgorithm2(b *testing.B) {
    for i := 0; i < b.N; i++ {
	Algorithm2(600851475143)
    }
}

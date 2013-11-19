package main

import (
	"testing"
)

func TestProblem1(t *testing.T) {
	result := problem1()
	if result != 233168 {
		t.Errorf("Expected answer: %v, obtained result: %v", 233168, result)
	}
}

package main

import "testing"

func TestProblem1V2(t *testing.T) {
	result := Problem1V2()
	if result != 233168 {
		t.Errorf("Exp answer: %v, Result: %v", 233168, result)
	}
}

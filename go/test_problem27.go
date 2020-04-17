package main

import (
	"testing"
)

func TestQuadraticPrime(t *testing.T) {
	testData := map[int]int{
		200:  -4925,
		500:  -18901,
		800:  -43835,
		1000: -59231,
	}
	for input, expected := range testData {
		actual := QuadraticPrime(input)
		if actual != expected {
			t.Errorf("TestQuadraticPrime(%d) failed - Want: %d, whereas Got: %d \n", input, expected, actual)
		}
	}
}

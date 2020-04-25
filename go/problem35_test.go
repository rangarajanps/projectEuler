package main

import (
	"testing"
)

func TestCircularPrimeCount(t *testing.T) {
	testData := map[int]int{
		100:     13,
		100000:  43,
		250000:  45,
		500000:  49,
		750000:  49,
		1000000: 55,
	}
	for input, expected := range testData {
		actual := CircularPrimeCount(input)
		if actual != expected {
			t.Errorf("TestCircularPrimeCount(%d) failed - Want: %d , whereas Got: %d \n", input, expected, actual)
		}
	}
}

package main

import (
	"testing"
)

func TestConsecutivePrimeSum(t *testing.T) {
	testData := map[int]int{
		100:     41,
		1000:    953,
		1000000: 997651,
	}

	for input, expected := range testData {
		actual := ConsecutivePrimeSum(input)
		if actual != expected {
			t.Errorf("TestConsecutivePrimeSum(%d) failed - Want: %d , wheras Got: %d \n", input, expected, actual)
		}
	}
}

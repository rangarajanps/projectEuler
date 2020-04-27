package main

import (
	"testing"
)

func TestTruncatablePrimeSum(t *testing.T) {

	testData := map[int]int{
		8:  1986,
		9:  5123,
		10: 8920,
		11: 748317,
	}
	for input, expected := range testData {
		actual := TruncatablePrimeSum(input)
		if actual != expected {
			t.Errorf("TestTruncatablePrimeSum(%d) failed - Want: %d , whereas Got: %d \n", input, expected, actual)
		}
	}

}

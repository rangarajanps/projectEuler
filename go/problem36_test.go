package main

import (
	"testing"
)

func TestDoubleBasePalindromeSum(t *testing.T) {
	testData := map[int]int{
		1000:    1772,
		50000:   105795,
		500000:  286602,
		1000000: 872187,
	}
	for input, expected := range testData {
		actual := DoubleBasePalindromeSum(input)
		if actual != expected {
			t.Errorf("TestDoubleBasePalindromeSum(%d) failed - Want: %d whereas, Got: %d \n", input, expected, actual)
		}
	}
}

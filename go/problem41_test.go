package main

import (
	"testing"
)

func TestLargestNLengthPandigitalPrime(t *testing.T) {

	testData := map[int]int{
		4: 4231,
		7: 7652413,
	}
	for input, expected := range testData {
		actual := LargestNLengthPandigitalPrime(input)
		if expected != actual {
			t.Errorf("TestLargestNLengthPandigitalPrime(%d) failed - Want: %d , whereas Got: %d \n", input, expected, actual)
		}
	}
}

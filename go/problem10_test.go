package main

import (
	"testing"
)

func TestFindSumOfPrimesTill(t *testing.T) {

	testData := map[int]int{
		17:      41,
		2001:    277050,
		140759:  873608362,
		2000000: 142913828922,
	}

	for input, expected := range testData {
		actualResult := FindSumOfPrimesTill(input)
		if actualResult != expected {
			t.Errorf("TestFindSumOfPrimesTill of (%d) was incorrect, got: %d, want: %d.", input, actualResult, expected)
		}
	}
}

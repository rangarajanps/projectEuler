package main

import (
	"testing"
)

func TestGenerateNthPrimeNumber(t *testing.T) {

	testcases := map[int]int{
		6:     13,
		10:    29,
		100:   541,
		1000:  7919,
		10001: 104743,
	}

	for input, expected := range testcases {
		actualResult := GenerateNthPrimeNumber(input)
		if actualResult != expected {
			t.Errorf("TestGenerateNthPrimeNumber of (%d) was incorrect, got: %d, want: %d.", input, actualResult, expected)
		}
	}
}

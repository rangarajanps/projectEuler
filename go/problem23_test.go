package main

import (
	"testing"
)

func TestNonAbundantSum(t *testing.T) {

	testData := map[int]int{
		10000: 3731004,
		15000: 4039939,
	}
	for input, expected := range testData {
		actual := NonAbundantSum(input)
		if actual != expected {
			t.Errorf("TestNonAbundantSum(%d) failed - Want: %d whereas, Got: %d \n", input, expected, actual)
		}
	}
}

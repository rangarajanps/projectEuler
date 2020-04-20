package main

import (
	"testing"
)

func TestSumDigitNPowers(t *testing.T) {
	testData := map[int]int{
		2: 0,
		3: 1301,
		4: 19316,
		5: 443839,
	}
	for input, expected := range testData {
		actual := SumDigitNPowers(input)
		if actual != expected {
			t.Errorf("TestSumDigitNPowers(%d) failed - Want: %d whereas, Got: %d \n", input, expected, actual)
		}
	}
}

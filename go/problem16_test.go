package main

import (
	"testing"
)

func TestPowerDigitSum(t *testing.T) {
	testData := map[int]int{
		15:   26,
		128:  166,
		1000: 1366,
	}

	for input, expected := range testData {
		actual := PowerDigitSum(input)
		if actual != expected {
			t.Errorf("TestPowerDigitSum(%d) failed - Want: %d , wheras Got: %d \n", input, expected, actual)
		}
	}
}

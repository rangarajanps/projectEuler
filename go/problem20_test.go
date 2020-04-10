package main

import (
	"testing"
)

func TestFactorialDigitSum(t *testing.T) {

	testData := map[int]int{
		10:  27,
		25:  72,
		50:  216,
		75:  432,
		100: 648,
	}
	for input, expected := range testData {
		actual := FactorialDigitSum(input)
		if actual != expected {
			t.Errorf("FactorialDigitSum(%d) failed - Want: %d , whereas Got: %d \n", input, expected, actual)
		}
	}
}

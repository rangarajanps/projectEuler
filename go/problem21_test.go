package main

import (
	"testing"
)

func TestFindAmicableNumberSum(t *testing.T) {

	testData := map[int]int{
		1000:  504,
		2000:  2898,
		5000:  8442,
		10000: 31626,
	}
	for input, expected := range testData {
		actual := FindAmicableNumberSum(input)
		if actual != expected {
			t.Errorf("FindAmicableNumberSum(%d) failed - Want: %d , whereas Got: %d \n", input, expected, actual)
		}
	}
}

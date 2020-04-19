package main

import (
	"testing"
)

func TestCountDistinctPowers(t *testing.T) {
	testData := map[int]int{
		15: 177,
		20: 324,
		25: 519,
		30: 755,
	}
	for input, expected := range testData {
		actual := CountDistinctPowers(input)
		if actual != expected {
			t.Errorf("TestCountDistinctPowers(%d) failed - Want: %d , whereas Got: %d \n", input, expected, actual)
		}
	}
}

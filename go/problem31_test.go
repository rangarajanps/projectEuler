package main

import (
	"testing"
)

func TestCoinsSum(t *testing.T) {
	testData := map[int]int{
		50:  451,
		100: 4563,
		150: 21873,
		200: 73682,
	}
	for input, expected := range testData {
		actual := CoinsSum(input)
		if actual != expected {
			t.Errorf("TestCoinsSum(%d) failed - Want: %d , whereas Got: %d \n", input, expected, actual)
		}
	}
}

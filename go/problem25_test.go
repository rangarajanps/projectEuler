package main

import (
	"testing"
)

func TestNthFiboNumber(t *testing.T) {
	testData := map[int]int{
		5:  21,
		10: 45,
		15: 69,
		20: 93,
	}
	for input, expected := range testData {
		actual := NthFiboNumber(input)
		if actual != expected {
			t.Errorf("TestNthFiboNumber(%d) failed - Want: %d, whereas Got: %d \n", input, expected, actual)
		}
	}
}

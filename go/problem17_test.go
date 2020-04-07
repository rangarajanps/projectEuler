package main

import (
	"testing"
)

func TestNumberLetterCount(t *testing.T) {

	testData := map[int]int{
		5: 19, 150: 1903, 1000: 21124,
	}
	for input, expected := range testData {
		actual := NumberLetterCount(input)
		if actual != expected {
			t.Errorf("TestNumberLetterCount(%d) failed - Want: %d , whereas Got: %d \n", input, expected, actual)
		}
	}
}

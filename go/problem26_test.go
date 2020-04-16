package main

import (
	"testing"
)

func TestReciprocalCycles(t *testing.T) {
	testData := map[int]int{
		700:  659,
		800:  743,
		900:  887,
		1000: 983,
	}
	for input, expected := range testData {
		actual := ReciprocalCycles(input)
		if actual != expected {
			t.Errorf("TestReciprocalCycles(%d) failed - Want: %d , whereas Got: %d \n", input, expected, actual)
		}
	}
}

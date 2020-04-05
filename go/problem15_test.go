package main

import (
	"testing"
)

func TestLatticePaths(t *testing.T) {

	testData := map[int]int{
		4:  70,
		9:  48620,
		20: 137846528820,
	}

	for input, expected := range testData {
		actual := LatticePaths(input)
		if actual != expected {
			t.Errorf("TestLatticePaths(%d) failed, got: %d , whereas want: %d \n", input, actual, expected)
		}
	}
}

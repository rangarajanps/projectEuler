package main

import (
	"testing"
)

func TestCountCodedTriangleNumbers(t *testing.T) {
	testData := map[int]int{
		1400: 129,
		1500: 137,
		1600: 141,
		1786: 162,
	}
	for input, expected := range testData {
		actual := CountCodedTriangleNumbers(input)
		if expected != actual {
			t.Errorf("TestCountCodedTriangleNumbers(%d) failed - Want: %d , whereas Got: %d \n", input, expected, actual)
		}
	}
}

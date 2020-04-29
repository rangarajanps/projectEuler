package main

import (
	"testing"
)

func TestIntRightTriangles(t *testing.T) {

	testData := map[int]int{
		500:  420,
		800:  720,
		900:  840,
		1000: 840,
	}
	for input, expected := range testData {
		actual := IntRightTriangles(input)
		if expected != actual {
			t.Errorf("TestIntRightTriangles(%d) failed - Want: %d , whereas Got: %d \n", input, expected, actual)
		}
	}

}

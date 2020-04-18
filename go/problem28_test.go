package main

import (
	"testing"
)

func TestSpiralDiagonalsSum(t *testing.T) {
	testData := map[int]int{
		101:  692101,
		303:  18591725,
		505:  85986601,
		1001: 669171001,
	}
	for input, expected := range testData {
		actual := SpiralDiagonalsSum(input)
		if actual != expected {
			t.Errorf("TestSpiralDiagonalsSum(%d) failed - Want: %d, whereas Got: %d \n", input, expected, actual)
		}
	}
}

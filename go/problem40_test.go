package main

import (
	"testing"
)

func TestChampernowneConstant(t *testing.T) {
	testData := map[int]int{
		100:     5,
		1000:    15,
		1000000: 210,
	}
	for input, expected := range testData {
		actual := ChampernowneConstant(input)
		if expected != actual {
			t.Errorf("TestChampernowneConstant(%d) failed - Want: %d , whereas Got: %d \n", input, expected, actual)
		}
	}

}

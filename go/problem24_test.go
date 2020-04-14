package main

import (
	"testing"
)

func TestLexicographicPermutation(t *testing.T) {
	testData := map[int]int{
		699999: 1938246570,
		899999: 2536987410,
		900000: 2537014689,
		999999: 2783915460,
	}
	for input, expected := range testData {
		actual := LexicographicPermutation(input)
		if actual != expected {
			t.Errorf("TestLexicographicPermutation(%d) failed - Want: %d , whereas Got: %d \n", input, expected, actual)
		}
	}
}

package main

import (
	"testing"
)

func TestFindLongestCollatzSequence(t *testing.T) {

	testData := map[int]int{
		14:     9,
		5847:   3711,
		46500:  35655,
		54512:  52527,
		100000: 77031,
	}
	for input, expected := range testData {
		actual := FindLongestCollatzSequence(input)
		if actual != expected {
			t.Errorf("TestFindLongestCollatzSequence(%d) failed, Want: %d, whereas Got: %d", input, expected, actual)
		}
	}
}

package main

import (
	"testing"
)

func TestFindLargestProduct(t *testing.T) {

	testData := map[int]int{
		4:  5832,
		13: 23514624000,
	}

	for input, expected := range testData {
		actualResult := FindLargestProduct(input)
		if actualResult != expected {
			t.Errorf("TestFindLargestProduct of (%d) was incorrect, got: %d, want: %d.", input, actualResult, expected)
		}
	}
}

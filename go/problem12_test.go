package main

import (
	"testing"
)

func TestFindHighlyDivisibleTriangularNumber(t *testing.T) {

	testData := map[int]int{
		5:   28,
		23:  630,
		167: 1385280,
		374: 17907120,
		500: 76576500,
	}

	for input, expected := range testData {
		actualResult, err := FindHighlyDivisibleTriangularNumber(input)
		if err != nil {
			t.Errorf("TestFindHighlyDivisibleTriangularNumber(%d) failed \n got: %v\n want: %d.", input, err, expected)
		} else if actualResult != expected {
			t.Errorf("TestFindHighlyDivisibleTriangularNumber of (%d) was incorrect, got: %d, want: %d.", input, actualResult, expected)
		}
	}

}

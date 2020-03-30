package main

import (
	"testing"
)

func testFindPythogoreanTriplet(t *testing.T) {

	testData := map[int][]int{
		24:   []int{480},
		120:  []int{49920, 55080, 60000},
		1000: []int{31875000},
	}

	for input, expected := range testData {
		actualResult, err := FindPythogoreanTriplet(input)
		if err != nil {
			t.Errorf("testFindPythogoreanTriplet of (%d) failed with error message (%v)", input, err)
		} else {
			var result bool = false
			for _, expValue := range expected {
				if actualResult == expValue {
					result = true
					break
				}
			}
			if !result {
				t.Errorf("testFindPythogoreanTriplet of (%d) was incorrect, got: %d, want: %d.", input, actualResult, expected)
			}
		}
	}
}

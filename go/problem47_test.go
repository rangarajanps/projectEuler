package main

import (
	"testing"
)

func TestDistinctPrimeFactors(t *testing.T) {
	list1 := []int{2, 2}
	list2 := []int{3, 3}
	list3 := []int{4, 4}
	input := [][]int{list1, list2, list3}
	expected := []int{14, 644, 134043}

	for i := 0; i < len(input); i++ {
		actual := DistinctPrimeFactors(input[i][0], input[i][1])
		if actual != expected[i] {
			t.Errorf("TestDistinctPrimeFactors(%d, %d) failed - Want: %d whereas, Got: %d \n", input[i][0], input[i][1], expected[i], actual)
		}
	}
}

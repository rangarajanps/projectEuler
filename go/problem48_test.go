package main

import (
	"testing"
)

func TestSelfPowers(t *testing.T) {
	list1 := []int{10, 3}
	list2 := []int{150, 5}
	list3 := []int{673, 7}
	list4 := []int{1000, 10}
	input := [][]int{list1, list2, list3, list4}
	expected := []string{"317", "29045", "2473989", "9110846700"}

	for i := 0; i < len(input); i++ {
		actual := SelfPowers(input[i][0], input[i][1])
		if actual != expected[i] {
			t.Errorf("TestSelfPowers(%d, %d) failed - Want: %v whereas, Got: %v \n", input[i][0], input[i][1], expected[i], actual)
		}
	}
}

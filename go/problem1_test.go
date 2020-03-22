package problem1

import (
	"testing"
)

func TestSumOfMultiplesOf3And5(t *testing.T) {

	tables := []struct {
		input    int
		expected int
	}{
		{10, 23},
		{49, 543},
		{1000, 233168},
		{8456, 16687353},
		{19564, 89301183},
	}

	for _, table := range tables {
		total := SumOfMultiplesOf3And5(table.input)
		if total != table.expected {
			t.Errorf("TestSumOfMultiplesOf3And5 of (%d) was incorrect, got: %d, want: %d.", table.input, total, table.expected)
		}
	}
}

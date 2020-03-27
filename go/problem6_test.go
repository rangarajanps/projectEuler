package main

import (
	"testing"
)

func TestFindSumSquareDifference(t *testing.T) {

	tables := []struct {
		input    int
		expected int
	}{
		{10, 2640},
		{20, 41230},
		{100, 25164150},
	}

	for _, table := range tables {
		total := FindSumSquareDifference(table.input)
		if total != table.expected {
			t.Errorf("TestFindSumSquareDifference of (%d) was incorrect, got: %d, want: %d.", table.input, total, table.expected)
		}
	}
}

package main

import (
	"testing"
)

func TestSmallestNumber(t *testing.T) {

	tables := []struct {
		input    int
		expected int
	}{
		{5, 60},
		{7, 420},
		{10, 2520},
		{13, 360360},
		{20, 232792560},
	}

	for _, table := range tables {
		total := SmallestNumber(table.input)
		if total != table.expected {
			t.Errorf("TestSmallestNumber of (%d) was incorrect, got: %d, want: %d.", table.input, total, table.expected)
		}
	}
}

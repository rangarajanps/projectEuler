package main

import (
	"testing"
)

func TestLargestPrimeFactor(t *testing.T) {

	tables := []struct {
		input    int
		expected int
	}{
		{2, 2},
		{3, 3},
		{5, 5},
		{7, 7},
		{13195, 29},
		{600851475143, 6857},
	}

	for _, table := range tables {
		total := LargestPrimeFactor(table.input)
		if total != table.expected {
			t.Errorf("TestLargestPrimeFactor of (%d) was incorrect, got: %d, want: %d.", table.input, total, table.expected)
		}
	}
}

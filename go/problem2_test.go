package main

import (
	"testing"
)

func TestSumOfEvenFibonacciNumber(t *testing.T) {

	tables := []struct {
		input    int
		expected int
	}{
		{10, 10},
		{60, 44},
		{1000, 798},
		{100000, 60696},
		{4000000, 4613732},
	}

	for _, table := range tables {
		total := SumOfEvenFibonacciNumber(table.input)
		if total != table.expected {
			t.Errorf("TestSumOfEvenFibonacciNumber of (%d) was incorrect, got: %d, want: %d.", table.input, total, table.expected)
		}
	}
}

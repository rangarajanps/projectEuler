package main

import (
	"testing"
)

func TestFindMaxPalindrome(t *testing.T) {

	tables := []struct {
		input    int
		expected int
	}{
		{2, 9009},
		{3, 906609},
	}

	for _, table := range tables {
		total := FindMaxPalindrome(table.input)
		if total != table.expected {
			t.Errorf("TestFindMaxPalindrome of (%d) was incorrect, got: %d, want: %d.", table.input, total, table.expected)
		}
	}
}

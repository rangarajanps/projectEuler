package main

import (
	"testing"
)

func TestLargestPandigitalMultiples(t *testing.T) {

	actual := LargestPandigitalMultiples()
	expected := 932718654
	if actual != expected {
		t.Errorf("TestLargestPandigitalMultiples() failed - Want: %d whereas, Got: %d \n", expected, actual)
	}
}

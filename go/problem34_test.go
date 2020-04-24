package main

import (
	"testing"
)

func TestDigitFactorialsSum(t *testing.T) {

	actual := DigitFactorialsSum()
	expected := 40730
	if actual != expected {
		t.Errorf("TestDigitFactorialsSum() failed - Want: %d whereas, Got: %d \n", expected, actual)
	}
}

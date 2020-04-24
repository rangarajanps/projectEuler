package main

import (
	"testing"
)

func TestLowestDenominatorOfDigitCancelling(t *testing.T) {

	actual := LowestDenominatorOfDigitCancelling()
	expected := 100
	if actual != expected {
		t.Errorf("TestLowestDenominatorOfDigitCancelling() failed - Want: %d whereas, Got: %d \n", expected, actual)
	}
}

package main

import (
	"testing"
)

func TestPentagonNumbers(t *testing.T) {

	actual := PentagonNumbers()
	expected := 5482660
	if actual != expected {
		t.Errorf("TestPentagonNumbers() failed - Want: %d whereas, Got: %d \n", expected, actual)
	}

}

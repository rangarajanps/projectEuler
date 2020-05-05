package main

import (
	"testing"
)

func TestTriPentaHexaNumbers(t *testing.T) {

	actual := TriPentaHexaNumbers()
	var expected float64 = 1533776805
	if actual != expected {
		t.Errorf("TestTriPentaHexaNumbers() failed - Want: %v whereas, Got: %v \n", expected, actual)
	}

}

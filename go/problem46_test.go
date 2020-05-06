package main

import (
	"testing"
)

func TestGoldbachConjecture(t *testing.T) {

	actual := GoldbachConjecture(6000)
	expected := 5777
	if actual != expected {
		t.Errorf("TestGoldbachConjecture() failed - Want: %d whereas, Got: %d \n", expected, actual)
	}

}

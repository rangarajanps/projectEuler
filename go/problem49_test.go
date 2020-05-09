package main

import (
	"testing"
)

func TestPrimePermutations(t *testing.T) {
	actual := PrimePermutations()
	if actual != "296962999629" {
		t.Errorf("TestPrimePermutations() failed - Want: 296962999629 , whereas Got: %v \n", actual)
	}
}

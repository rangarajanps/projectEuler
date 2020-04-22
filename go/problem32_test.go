package main

import (
	"testing"
)

func TestPandigitProducts(t *testing.T) {
	actual := PandigitProducts()
	if actual != 45228 {
		t.Errorf("TestPandigitProducts() failed - Want: 45228 , whereas Got: %d \n", actual)
	}
}

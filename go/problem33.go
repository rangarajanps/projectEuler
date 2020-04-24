package main

import (
	"fmt"
)

func main() {
	fmt.Println(LowestDenominatorOfDigitCancelling())
}

func LowestDenominatorOfDigitCancelling() int {

	productNum := 1
	productDen := 1

	for a := 10; a < 100; a++ {
		for b := a + 1; b < 100; b++ {
			if isDigitCancelling(a, b) {
				productNum *= a
				productDen *= b
			}
		}
	}

	greatestFactor := gcf(productNum, productDen)

	return productDen / greatestFactor
}

// Function that returns whether a fraction is a digit cancelling fraction a/b
func isDigitCancelling(a int, b int) bool {

	// If they're both multiples of ten, return false
	if a%10 == 0 && b%10 == 0 {
		return false
	}

	// Get all the relevant digits isolated
	numTens := a / 10
	numOnes := a % 10
	denTens := b / 10
	denOnes := b % 10

	// Check all cases
	// numTens = denTens
	if numTens == denTens && isSameFraction(a, b, numOnes, denOnes) {
		return true
	}
	// numTens = denOnes
	if numTens == denOnes && isSameFraction(a, b, numOnes, denTens) {
		return true
	}
	// numOnes = denTens
	if numOnes == denTens && isSameFraction(a, b, numTens, denOnes) {
		return true
	}
	// numOnes = denOnes
	if numOnes == denOnes && isSameFraction(a, b, numTens, denTens) {
		return true
	}

	return false
}

// Function that checks if two fractions a/b and c/d are the same
func isSameFraction(a, b, c, d int) bool {
	return a*d == b*c
}

// Function that finds the GCF of two numbers
func gcf(a int, b int) int {
	smallest := a
	if a > b {
		smallest = b
	}

	gcFactor := 1

	for i := 1; i <= smallest; i++ {
		if a%i == 0 && b%i == 0 {
			gcFactor = i
		}
	}

	return gcFactor
}

// Function to return the lowest denominator of fraction a/b
func lowestDenominator(a int, b int) int {
	// First find gcf
	greatestFactor := gcf(a, b)

	return b / greatestFactor
}

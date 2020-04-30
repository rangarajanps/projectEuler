package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(ChampernowneConstant(100))
}

func ChampernowneConstant(n int) int {

	magnitude := 1
	product := 1

	for i := 0; magnitude <= n; i++ {
		temp := nthDigitOfFractional(magnitude)
		product *= temp
		magnitude *= 10
	}

	return product
}

// Function to find the nth digit (going from right to left) of a number p
func nthDigitOfInteger(n int, p int) int {

	for n > 0 {
		p /= 10
		n--
	}
	return p % 10
}

// Function to find the nth digit of the fractional concatenation of the positive integers
func nthDigitOfFractional(n int) int {

	currPlace := 1
	currNumDigits := 1

	for n > 0 {
		for i := currPlace; i < currPlace*10; i++ {
			n -= currNumDigits
			if n <= 0 {
				return nthDigitOfInteger(int(math.Abs(float64(n))), i)
			}
		}
		currNumDigits++
		currPlace *= 10
	}

	return 0
}

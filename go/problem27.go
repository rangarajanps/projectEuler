package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(QuadraticPrime(100))
}

func QuadraticPrime(limit int) int {

	maxA, maxB, maxN := 0, 0, 0

	for a := -limit; a <= limit; a++ {

		for b := -limit; b <= limit; b++ {

			count := 0
			for count < b {
				if isPrime(quadraticVal(a, b, count)) {
					count++
				} else {
					break
				}
			}

			if count > maxN {
				maxN = count
				maxA = a
				maxB = b
			}
		}
	}

	return maxA * maxB

}

func quadraticVal(a, b, n int) int {
	return int(math.Abs(float64(n*n + a*n + b)))
}

func isPrime(number int) bool {

	if number <= 1 {
		return false
	}

	if number%2 == 0 {
		return false
	}

	for i := 3; i <= int(math.Sqrt(float64(number))); i += 2 {
		if number%i == 0 {
			return false
		}
	}

	return true
}

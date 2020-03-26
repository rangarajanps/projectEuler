package main

import (
	"fmt"
)

func main() {
	fmt.Println(SmallestNumber(13))
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func SmallestNumber(limit int) int {
	result := 1
	for j := 2; j <= limit; j++ {
		result = LCM(result, j)
	}

	return result
}

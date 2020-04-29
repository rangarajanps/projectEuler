package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(IntRightTriangles(500))
}

func IntRightTriangles(n int) int {

	maxPeri := 0
	maxCount := 0
	temp := 0

	for i := 8; i <= n; i++ {
		temp = countRightTrianglesWithPeri(i)
		if temp >= maxCount {
			maxCount = temp
			maxPeri = i
		}
	}

	return maxPeri
}

func countRightTrianglesWithPeri(n int) int {

	count := 0
	c := 0
	for a := 11; a < n/3; a++ {
		for b := a + 1; b < n/2; b++ {

			temp, frac := math.Modf(math.Sqrt(float64(a*a + b*b)))
			if frac != 0 {
				continue
			}
			c = int(temp)
			if a+b+c == n {
				count++
				// fmt.Printf("Found a right triangle for Perimeter %d with sides a=%d b=%d c=%d \n", n, a, b, c)
			}
		}
	}
	return count
}

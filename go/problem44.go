package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(PentagonNumbers())
}

func isPentagonal(n int) bool {
	p := (1 + math.Sqrt(1+24*float64(n))) / 6
	return p == math.Floor(p)
}

func solvePentagonal(n int) int {
	return (n * (3*n - 1)) / 2
}

func PentagonNumbers() int {

	found := false
	i := 1
	var ans int

	for !found {
		p1 := solvePentagonal(i)
		for j := i - 1; j > 0; j-- {
			p2 := solvePentagonal(j)
			if isPentagonal(p1+p2) && isPentagonal(p1-p2) {
				found = true
				ans = p1 - p2
				break
			}
		}
		i++
	}

	return ans
}

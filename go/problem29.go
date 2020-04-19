package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(CountDistinctPowers(5))
}

func CountDistinctPowers(limit int) int {

	powerSlice := make(map[float64]float64, 10)
	for i := 2.0; i <= float64(limit); i++ {
		for j := 2.0; j <= float64(limit); j++ {
			powerSlice[math.Pow(i, j)] += 1
		}
	}

	return len(powerSlice)
}

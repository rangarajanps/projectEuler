package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var value float64 = TriPentaHexaNumbers()
	result := strconv.FormatFloat(value, 'f', 0, 64)
	fmt.Println(result)
}

func isPentagonal(n float64) bool {
	p := (1 + math.Sqrt(1+24*float64(n))) / 6
	return p == math.Floor(p)
}

func isHexagonal(n float64) bool {
	p := (1 + math.Sqrt(1+8*float64(n))) / 6
	return p == math.Floor(p)
}

func TriPentaHexaNumbers() float64 {

	var numToTest float64 = 285
	var result float64 = 0
	for true {

		numToTest++
		result = numToTest * (2*numToTest - 1)

		if isPentagonal(result) && isHexagonal(result) {
			return result
		}

	}
	return -1
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println(FindSumSquareDifference(10))
}

func FindSumSquareDifference(limit int) int {

	sum := (limit * (limit + 1)) / 2
	sumSquare := sum * sum
	squareSum := 0
	for i := 1; i <= limit; i++ {
		squareSum += i * i
	}
	return sumSquare - squareSum

}

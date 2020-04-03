package main

import (
	"fmt"
	"math"
)

func main() {
	v, err := FindHighlyDivisibleTriangularNumber(5)
	if err != nil {
		fmt.Println("Error : ", err)
	}
	fmt.Println(v)
}

func divisorCount(num int) int {
	
	incr := 2
	count := 2
	if num%2 == 0 {
		incr = 1
		count = 4
	}

	for i := 3; i <= int(math.Trunc(math.Sqrt(float64(num)))); i += incr {
		if num%i == 0 {
			count += 2
		}
	}

	return count
}

func FindHighlyDivisibleTriangularNumber(limit int) (int, error) {

	triNum := 1
	incr := 2
	for true {
		triNum = triNum + incr
		//fmt.Printf(" %v ", triNum)
		if divisorCount(triNum) >= limit {
			return triNum, nil
		}
		incr++
	}

	return -1, fmt.Errorf("Could not find a triangular number with %v divisors", limit)
}

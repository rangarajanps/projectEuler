package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(FindAmicableNumberSum(1000))
}

func FindAmicableNumberSum(limit int) int {

	factorSumMapSlice := make(map[int]int)
	for i := 2; i <= limit; i++ {
		factorSumMapSlice[i] = factorSum(i)
	}

	amicableNumberSum := 0
	for index, val := range factorSumMapSlice {
		if index == factorSumMapSlice[val] && index != val {
			// fmt.Println("Found an amicable number ", val, index, factorSumMapSlice[val])
			amicableNumberSum = amicableNumberSum + val + index
			factorSumMapSlice[val] = -1 // to avoid duplicate count or need to divide by 2
		}
	}

	return amicableNumberSum
}

func factorSum(number int) int {

	var factorSum int = 1
	var i int
	var incrementor int = 1
	if number%2 != 0 {
		incrementor = 2
	} else {
		quotient := number / 2
		if quotient != 2 {
			factorSum += quotient
		}
		factorSum += 2
	}

	for i = 3; i <= int(math.Sqrt(float64(number))); i += incrementor {
		if number%i == 0 {
			quotient := number / i
			if quotient != i {
				factorSum += quotient
			}
			factorSum += i
		}
	}

	return factorSum
}

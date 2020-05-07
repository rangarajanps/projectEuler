package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(DistinctPrimeFactors(2, 2))
}

func DistinctPrimeFactors(targetFactorNum int, targetConsecutiveCount int) int {

	consecutiveNumCount := 0
	factorCount := 0
	startingNumber := 0
	numToCheck := 10
	foundOne := false

	for true {

		if isPrime(numToCheck) {
			numToCheck++
			foundOne = false
			consecutiveNumCount = 0
			continue
		}

		factorCount = calculatePrimeFactors(numToCheck)
		if factorCount == targetFactorNum {

			if !foundOne {
				// fmt.Println("Initializing consecutive count to 1 for ", numToCheck);
				foundOne = true
				startingNumber = numToCheck
				consecutiveNumCount = 1
			} else {
				// fmt.Println("Incrementing consecutive count by 1 for " + numToCheck);
				consecutiveNumCount++
				if consecutiveNumCount == targetConsecutiveCount {
					// fmt.Println("Returning for " + numToCheck);
					return startingNumber
				}
			}
		} else {
			foundOne = false
			consecutiveNumCount = 0
		}
		numToCheck++
	}

	return -1

}

func calculatePrimeFactors(num int) int {

	factors := make([]int, 0, 10)
	factorsCountMap := make(map[int]int)

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {

		if num%i == 0 {
			numToTest := num / i
			if isPrime(i) {
				updateMap(factorsCountMap, i)
			}
			if isPrime(numToTest) {
				updateMap(factorsCountMap, numToTest)
			}
			if numToTest == i {
				continue
			}

			for numToTest > 0 {
				if numToTest%i == 0 {
					if isPrime(i) {
						updateMap(factorsCountMap, i)
					}
					numToTest = numToTest / i
				} else {
					break
				}
			}
		}
	}

	for key, value := range factorsCountMap {
		if value == 1 {
			factors = append(factors, key)
		} else {
			product := 1
			for p := 0; p < value; p++ {
				product *= value
			}
			factors = append(factors, product)
		}
	}

	return len(factors)
}

func updateMap(fMap map[int]int, factor int) map[int]int {

	value, ok := fMap[factor]
	if !ok {
		fMap[factor] = 1
	} else {
		fMap[factor] = value + 1
	}
	return fMap
}

func isPrime(num int) bool {

	if num <= 1 {
		return false
	}

	if num == 2 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	for i := 3; i <= int(math.Sqrt(float64(num))); i += 2 {
		if num%i == 0 {
			return false
		}
	}

	return true
}

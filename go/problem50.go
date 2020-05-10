package main

import (
	"fmt"
)

func main() {
	fmt.Println(ConsecutivePrimeSum(1000))
}

func ConsecutivePrimeSum(uptoN int) int {

	isPrime := make(map[int]bool)
	for i := 2; i <= uptoN; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i <= uptoN; i++ {
		if isPrime[i] {
			for j := 2; i*j <= uptoN; j++ {
				isPrime[i*j] = false
			}
		}
	}

	primeNumbers := make([]int, 0, 100)
	for i := 2; i <= len(isPrime); i++ {
		if isPrime[i] {
			primeNumbers = append(primeNumbers, i)
		}
	}

	finalSum := 0
	finalCount := 0

	for start := 0; start < len(primeNumbers); start++ {
		sum := 0
		count := 0

		for current := start; current < len(primeNumbers); current++ {

			sum += primeNumbers[current]
			count++
			if sum >= uptoN {
				break
			}
			if isPrime[sum] && count > finalCount {
				finalCount = count
				finalSum = sum
			}
		}
	}

	return finalSum
}

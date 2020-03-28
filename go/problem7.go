package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(GenerateNthPrimeNumber(10))
}

func GenerateNthPrimeNumber(limit int) int {

	count := 1
	n := 3
	for {
		if isPrime(n) {
			count++
			if count == limit {
				break
			}
		}
		n += 2
	}
	return n
}

func isPrime(num int) bool {

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

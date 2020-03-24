package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(LargestPrimeFactor(13195))
}

func Sieve(n int) []int {
	var primes = make([]bool, n+1)
	for i := 1; i <= n; i++ {
		primes[i] = true
	}
	for i := 2; i < n/2; i++ {
		for j := 2; i*j < n; j++ {
			primes[i*j] = false
		}
	}
	var primeNumbers []int
	for i := 1; i <= n; i++ {
		if primes[i] {
			primeNumbers = append(primeNumbers, i)
		}
	}

	return primeNumbers
}

func LargestPrimeFactor(n int) int {
	
	if isPrime(n) {
		return n
	}
  
  var primeList = Sieve(int(math.Sqrt(float64(n))))
	var largestPrimeFactor = 0
	for i := 0; i < len(primeList); i++ {
		if n%primeList[i] == 0 && primeList[i] > largestPrimeFactor {
			largestPrimeFactor = primeList[i]
		}
	}
	return largestPrimeFactor
}

func isPrime(n int) bool {

	if n == 1 {
		return false
	}
	for x := 2; x < int(math.Sqrt(float64(n))); x++ {
		if n%x == 0 {
			return false
		}
	}

	return true
}

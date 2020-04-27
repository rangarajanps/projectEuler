package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(TruncatablePrimeSum(8))
}

func TruncatablePrimeSum(requiredNum int) int {

	count := 0
	sum := 0
	num := 11
	for count < requiredNum {
		if isPrime(num) && isTruncatablePrime(num) {
			count++
			sum += num
			// fmt.Println("Found a truncatable prime ", num)
		}
		num += 2
	}
	return sum
}

func isTruncatablePrime(num int) bool {

	for div := 10; div < num; div *= 10 {
		if !isPrime(int(num/div)) || !isPrime(num%div) {
			return false
		}
	}
	return true
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

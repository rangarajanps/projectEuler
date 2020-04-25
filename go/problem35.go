package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(CircularPrimeCount(100))
}

func CircularPrimeCount(limit int) int {

	count := 5
	// 2, 3, 5, 7, 11 is added as Circular Prime so that loop can be started
	// from 13 to skip all single digit numbers and rotated number being same as
	// initial input number. And also check only for odd numbers

	for i := 13; i < limit; i += 2 {

		if isPrime(i) {

			number := i
			digits := numberOfDigits(number)
			powTen := int(math.Pow10(digits - 1))

			isCircularPrime := true
			for i := 0; i < digits-1; i++ {
				rem := number % 10
				quo := number / 10
				number = rem*powTen + quo
				if !isPrime(number) {
					isCircularPrime = false
					break
				}
			}

			if isCircularPrime {
				// fmt.Println("Found a circular Prime ", number)
				count++
			}
		}

	}
	return count

}

// Function to return the count of digits of n
func numberOfDigits(n int) int {

	cnt := 0
	for n > 0 {
		cnt++
		n /= 10
	}
	return cnt
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

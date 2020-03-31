package main

import (
	"fmt"
)

func main() {
	fmt.Println(FindSumOfPrimesTill(17))
}

func seive(limit int) []int {
	var isPrimeList = make([]bool, limit+1)
	for i := 0; i < limit; i++ {
		isPrimeList[i] = true
	}
	isPrimeList[0] = false
	isPrimeList[1] = false

	for i := 2; i <= (limit/2)+1; i++ {
		for j := 2; i*j <= limit; j++ {
			isPrimeList[i*j] = false
		}
	}

	var primeList []int

	for i := 0; i < limit; i++ {
		if isPrimeList[i] {
			primeList = append(primeList, i)
		}
	}

	return primeList
}

func FindSumOfPrimesTill(limit int) int {

	primeNumbers := seive(limit)
	//fmt.Println(primeNumbers)
	sum := 0
	for _, value := range primeNumbers {
		sum += value
	}
	return sum

}

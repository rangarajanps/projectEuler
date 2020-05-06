package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(GoldbachConjecture(6000))
}

func GoldbachConjecture(limit int) int {

	composites, primes := sieveComposite(limit)

	for i := 0; i < len(composites); i++ {
		if composites[i]%2 == 0 {
			continue // skip even composite numbers
		}
		primeSubList := primes[:indexOf(composites[i], primes)]
		if !isGoldbach(composites[i], primeSubList) {
			return composites[i]
		}
	}

	return -1 // Not found
}

// Primes less than a given composite number is constructed
// element is the composite to check
func indexOf(element int, intList []int) int {
	for k, v := range intList {
		if v > element {
			return k
		}
	}
	return len(intList) //For not found, return the whole list
}

func isGoldbach(num int, primeList []int) bool {

	temp := 0
	for i := 0; i < len(primeList); i++ {
		temp = num - primeList[i]
		if temp%2 != 0 {
			continue
		}
		sqrtVal := math.Sqrt(float64(temp / 2))
		if sqrtVal == math.Trunc(sqrtVal) {
			return true
		}
	}
	return false

}

func sieveComposite(limit int) ([]int, []int) {

	if limit <= 1 {
		return []int{}, []int{}
	}

	isPrime := make(map[int]bool)
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}

	for i := 2; i <= limit/2; i++ {
		for j := 2; i*j <= limit; j++ {
			isPrime[i*j] = false
		}
	}

	var compNumbers []int
	var primeNumbers []int
	for i := 2; i <= limit; i++ {
		if !isPrime[i] {
			compNumbers = append(compNumbers, i)
		} else {
			primeNumbers = append(primeNumbers, i)
		}
	}

	return compNumbers, primeNumbers
}

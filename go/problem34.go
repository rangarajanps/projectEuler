package main

import (
	"fmt"
)

func main() {
	fmt.Println(DigitFactorialsSum())
}

var factorials = generateFactorials(9)

func DigitFactorialsSum() int {

	result := 0
	for n := 100; n <= 40888; n++ {
		if isDigitFactorial(n) {
			result += n
		}
	}
	return result
}

func isDigitFactorial(num int) bool {

	numToTest := num
	sum := 0
	rem := 0

	for numToTest >= 1 {
		rem = numToTest % 10
		if sum > num || factorials[rem] >= num {
			return false
		}

		sum = sum + factorials[rem]
		numToTest = numToTest / 10
	}

	if sum == num {
		return true
	}
	return false

}

func generateFactorials(limit int) map[int]int {

	factMap := make(map[int]int)
	factMap[0] = 1
	factMap[1] = 1
	factMap[2] = 2
	fact := 2
	for i := 3; i <= limit; i++ {
		fact = fact * i
		factMap[i] = fact
	}
	return factMap
}

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(NonAbundantSum(10000))
}

func NonAbundantSum(limit int) int {

	var abundantNums []int
	for i := 3; i <= limit; i++ {
		if IsAbundant(i) {
			abundantNums = append(abundantNums, i)
		}
	}
	// fmt.Println(len(abundantNums))

	nonAbundantNumbersSum := 0

	isSumOfAbundant := false

	for testNum := 1; testNum <= limit; testNum++ {

		isSumOfAbundant = false

		j := len(abundantNums) - 1
		for i := 0; abundantNums[i] < testNum; i++ {

			abundantNumber1 := abundantNums[i]
			if abundantNumber1 > testNum {
				break
			}

			abundantNumber2 := abundantNums[j]
			for (j > 0) && (abundantNumber1+abundantNumber2 > testNum) {
				j = j - 1
				abundantNumber2 = abundantNums[j]
			}

			if testNum == abundantNumber1+abundantNumber2 {
				// fmt.Println("Identified one ", testNum, " ", abundantNums[i], " ", abundantNums[j])
				isSumOfAbundant = true
				break
			}
		}

		if !isSumOfAbundant {
			nonAbundantNumbersSum += testNum
		}

	}

	return nonAbundantNumbersSum

}

func IsAbundant(number int) bool {

	if number <= 0 {
		return false
	}

	if number == 1 || number == 2 {
		return false
	}

	var factorSum int = 1
	var i int
	var incrementor int = 1
	if number%2 != 0 {
		incrementor = 2
	} else {
		quotient := int(number) / 2
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

	if factorSum > number {
		return true
	} else {
		return false
	}
}

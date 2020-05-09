package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(PrimePermutations())
}

func PrimePermutations() string {

	const UNUSUAL_CONST = 3330

	p1 := 1488

	for true {

		p2 := p1 + UNUSUAL_CONST
		p3 := p2 + UNUSUAL_CONST

		if isPrime(p1) && isPrime(p2) && isPrime(p3) {

			if checkIfThreeNumbersAreFormedOfSameDigit(p1, p2, p3) {
				var sb strings.Builder
				sb.WriteString(strconv.Itoa(p1))
				sb.WriteString(strconv.Itoa(p2))
				sb.WriteString(strconv.Itoa(p3))
				return sb.String()
			}
		}
		p1++
	}

	return ""
}

func checkIfThreeNumbersAreFormedOfSameDigit(n1 int, n2 int, n3 int) bool {
	num1AsStr := convertToString(n1)
	num2AsStr := convertToString(n2)
	num3AsStr := convertToString(n3)
	if strings.EqualFold(num1AsStr, num2AsStr) &&
		strings.EqualFold(num3AsStr, num2AsStr) &&
		strings.EqualFold(num1AsStr, num3AsStr) {
		return true
	}

	return false
}

func convertToString(num int) string {
	digits := make([]int, 0, 4)
	for num > 0 {
		digits = append(digits, num%10)
		num = num / 10
	}
	sort.Ints(digits)
	var sb strings.Builder
	for i := 0; i < len(digits); i++ {
		sb.WriteString(strconv.Itoa(digits[i]))
	}
	return sb.String()
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

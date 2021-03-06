package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(LargestNLengthPandigitalPrime(4))
}

func LargestNLengthPandigitalPrime(n int) int {

	var startStr strings.Builder
	for i := 1; i <= n; i++ {
		temp := strconv.Itoa(i)
		startStr.WriteString(temp)
	}
	startIndex, err := strconv.Atoi(startStr.String())
	if err != nil {
		fmt.Println("Error in start string to int conversion ", err)
		return -1
	}
	
	endIndex := 0
	orig_number := startIndex
	for orig_number > 0 {
		endIndex = endIndex*10 + orig_number%10
		orig_number /= 10
	}

	maxNLengthPandigitNum := 1
	for i := startIndex; i < endIndex; i++ {

		if isPandigital(i, n) && isPrime(i) {
			if i > maxNLengthPandigitNum {
				maxNLengthPandigitNum = i
			}
		}
	}
	return maxNLengthPandigitNum
}

func isPandigital(num int, uptoN int) bool {

	numAsStr := strconv.Itoa(num)

	digits := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := 0; i < uptoN; i++ {
		if strings.Index(numAsStr, digits[i]) == -1 {
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

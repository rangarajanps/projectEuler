package main

import (
	"fmt"
)

func main() {
	fmt.Println(DoubleBasePalindromeSum(1000))
}

func DoubleBasePalindromeSum(limit int) int {

	sum := 0
	for i := 1; i <= limit; i++ {
		if isPalindrome(i) && isBinaryPalindrome(i) {
			// fmt.Println("Found a Double base palindrome ", i)
			sum += i
		}
	}
	return sum
}

func isPalindrome(n int) bool {

	origNumber := n
	reversedNum := 0
	for n > 0 {
		reversedNum = reversedNum*10 + n%10
		n = n / 10
	}
	return reversedNum == origNumber

}

func isBinaryPalindrome(n int) bool {

	var numInBase2 []int
	itr := 0
	for n > 0 {
		numInBase2 = append(numInBase2, n%2)
		n = n / 2
		itr++
	}

	for i := 0; i < len(numInBase2)/2; i++ {
		if numInBase2[i] != numInBase2[len(numInBase2)-i-1] {
			return false
		}
	}

	return true

}

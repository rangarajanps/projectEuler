package main

import (
	"fmt"
)

func main() {
	fmt.Println(FactorialDigitSum(10))
}

func FactorialDigitSum(n int) int {

	if n <= 10 {
		return findSum(factorial(n))
	}

	fact := convertNumToArray(factorial(10))

	temp := 0
	carryFwd := 0
	for i := 11; i <= n; i++ {
		for j := 0; j < len(fact); j++ {
			temp = carryFwd + fact[j]*i
			fact[j] = temp % 10
			carryFwd = temp / 10
		}

		var carryFwdInArr []int
		if carryFwd > 0 {
			carryFwdInArr = convertNumToArray(carryFwd)
			for x := 0; x < len(carryFwdInArr); x++ {
				fact = append(fact, carryFwdInArr[x])
			}
			carryFwd = 0
		}

	}

	ans := 0
	for i := 0; i < len(fact); i++ {
		ans += fact[i]
	}
	return ans
}

func factorial(n int) int {

	fact := 2
	for i := 3; i <= n; i++ {
		fact = fact * i
	}
	return fact
}

func findSum(n int) int {

	sum := 0
	for n > 0 {
		sum = sum + n%10
		n = n / 10
	}
	return sum

}

func convertNumToArray(n int) []int {

	var num []int
	for n > 0 {
		num = append(num, n%10)
		n = n / 10
	}
	return num

}

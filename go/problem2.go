package main

import (
	"fmt"
)

func main() {

	fmt.Println(SumOfEvenFibonacciNumber(10))
}

func SumOfEvenFibonacciNumber(limit int) int {

	var fibSeries = make([]int, limit)
	fibSeries[0] = 1
	fibSeries[1] = 2

	for counter := 2; counter < limit; counter++ {
		if fibSeries[counter] > limit {
			break
		}
		fibSeries[counter] = fibSeries[counter-1] + fibSeries[counter-2]
	}
	//fmt.Println("Value of Fibonacci Sereis is ", fibSeries)

	sum := 0
	for j := 0; j < len(fibSeries); j++ {
		if fibSeries[j] > limit {
			break
		}
		if fibSeries[j]%2 == 0 {
			sum = sum + fibSeries[j]
		}
	}
	//fmt.Println("Sum of Even numbers is ", sum)
	return sum
}

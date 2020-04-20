package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(SumDigitNPowers(3))
}

func SumDigitNPowers(n int) int {

	result := 0
	init := strings.Repeat("1", n-1)
	startIndex, err1 := strconv.Atoi(init)
	if err1 != nil {
		return -1
	}
	if startIndex == 1 {
		startIndex++
	}
	end := strings.Repeat("9", n+1)
	endIndex, err2 := strconv.Atoi(end)
	if err2 != nil {
		return -1
	}
	for i := startIndex; i <= endIndex; i++ {
		if i == sumPower(i, float64(n)) {
			result += i
		}
	}
	return result
}

func sumPower(num int, n float64) int {
	sum := 0.0
	for num > 0 {
		sum += math.Pow(float64(num%10), n)
		num = num / 10
	}
	return int(sum)
}

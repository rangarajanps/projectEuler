package main

import (
	"fmt"
	"math/big"
)

func main() {
	results := SubDivisiblePandigitNumbers()
	for _, bigint := range results {
		fmt.Printf(bigint.String())
		fmt.Printf(" ")
	}
	fmt.Println()
}

func SubDivisiblePandigitNumbers() []big.Int {

	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	permuts := permutation(nums)
	result := make([][]int, 0, 10)

	for _, numSlice := range permuts {

		if isSubDivisible(numSlice) {
			result = append(result, numSlice)
		}
	}
	
	bigints := make([]big.Int, 0, len(result))
	for _, numSlice := range result {
		value := big.NewInt(0)
		for _, num := range numSlice {
			bigin := big.NewInt(int64(num))
			value = value.Mul(value, big.NewInt(10))
			value = value.Add(bigin, value)
		}
		bigints = append(bigints, *value)
	}
	return bigints
}

func permutation(xs []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}

func isSubDivisible(numbers []int) bool {

	// d2d3d4 = 406 is divisible by 2
	// d3d4d5 = 063 is divisible by 3
	// d4d5d6 = 635 is divisible by 5
	// d5d6d7 = 357 is divisible by 7
	// d6d7d8 = 572 is divisible by 11
	// d7d8d9 = 728 is divisible by 13
	// d8d9d10 = 289 is divisible by 17
	d2d3d4 := numbers[1]*100 + numbers[2]*10 + numbers[3]
	if d2d3d4%2 != 0 {
		return false
	}
	d3d4d5 := numbers[2]*100 + numbers[3]*10 + numbers[4]
	if d3d4d5%3 != 0 {
		return false
	}
	d4d5d6 := numbers[3]*100 + numbers[4]*10 + numbers[5]
	if d4d5d6%5 != 0 {
		return false
	}
	d5d6d7 := numbers[4]*100 + numbers[5]*10 + numbers[6]
	if d5d6d7%7 != 0 {
		return false
	}
	d6d7d8 := numbers[5]*100 + numbers[6]*10 + numbers[7]
	if d6d7d8%11 != 0 {
		return false
	}
	d7d8d9 := numbers[6]*100 + numbers[7]*10 + numbers[8]
	if d7d8d9%13 != 0 {
		return false
	}
	d8d9d10 := numbers[7]*100 + numbers[8]*10 + numbers[9]
	if d8d9d10%17 != 0 {
		return false
	}

	return true

}

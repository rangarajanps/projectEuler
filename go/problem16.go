package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	fmt.Println(PowerDigitSum(15))
}

func PowerDigitSum(exp int) int {

	n := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(exp)), nil).String()

	return sumOfDigits(n)
}

func sumOfDigits(num string) int {
	sum := 0
	for _, c := range num {
		i, err := strconv.Atoi(string(c))
		if err != nil {
			panic(err)
		}
		sum += i
	}
	return sum
}

package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(SelfPowers(10, 3))
}

func SelfPowers(uptoN int, numOfDigits int) string {

	sum := big.NewInt(1)
	num := big.NewInt(2)
	for i := 2; i <= uptoN; i++ {
		n := new(big.Int).Exp(big.NewInt(int64(i)), big.NewInt(int64(i)), nil)
		sum = sum.Add(sum, n)
		num = num.Add(num, big.NewInt(1))
	}
	value := sum.String()
	return value[len(value)-numOfDigits:]
}

package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("5th Fibo number is ", NthFiboNumber(5))
}

func NthFiboNumber(n int) int {

	if n == 1 {
		return 1
	}

	one := big.NewInt(int64(1))
	fiboSlice := map[int]*big.Int{
		0: one,
		1: one,
	}

	i := 2
	for true {

		var num = new(big.Int)
		num.Add(fiboSlice[i-1], fiboSlice[i-2])
		temp := num.String()
		if len(temp) == n {
			i += 1
			return i
		}
		fiboSlice[i] = num
		i += 1
	}

	return -1
}

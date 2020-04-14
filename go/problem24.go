package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(LexicographicPermutation(699999))
}

func factorial(num int) int {

	if num < 0 {
		return 0
	}

	product := 1
	for i := 2; i <= num; i++ {
		product *= i
	}
	return product
}

func LexicographicPermutation(remain int) int {

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	N := len(numbers)
	permNum := ""

	for i := 1; i < N; i++ {

		fact := factorial(N - i)
		j := remain / fact
		remain = remain % fact

		permNum = permNum + strconv.Itoa(numbers[j])
		numbers = append(numbers[:j], numbers[j+1:]...)

		if remain == 0 {
			break
		}
	}

	for i := 0; i < len(numbers); i++ {
		permNum = permNum + strconv.Itoa(numbers[i])
	}

	val, err := strconv.Atoi(permNum)
	if err != nil {
		fmt.Println("Error in converting string to int for ", permNum)
	}

	return val
}

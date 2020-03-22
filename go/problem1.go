package problem1

import (
	"fmt"
)

func main() {

	fmt.Println(SumOfMultiplesOf3And5(10))
}

func SumOfMultiplesOf3And5(limit int) int {

	result := 0
	for i := 1; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			result += i
		}
	}

	return result
}

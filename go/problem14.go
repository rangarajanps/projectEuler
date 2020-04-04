package main

import (
	"fmt"
)

func main() {
	fmt.Println(FindLongestCollatzSequence(14))
}

func genCollatzSequence(n int) int {
	if n%2 == 0 {
		return n / 2
	} else {
		return 3*n + 1
	}
}

func FindLongestCollatzSequence(limit int) int {

	var maxCount, maxNumber int

	for number := 2; number <= limit; number++ {

		result := countCollatzSequenceToOne(number)
		if result > maxCount {
			maxCount = result
			maxNumber = number
		}

	}

	return maxNumber
}

func countCollatzSequenceToOne(num int) int {

	count := 1
	// fmt.Printf("genCollatzSequence is ")
	for true {

		num = genCollatzSequence(num)
		count++
		// fmt.Printf("%d ", num)
		if num == 1 {
			break
		}

	}
	// fmt.Printf("\nCount is %d \n", count)
	return count

}

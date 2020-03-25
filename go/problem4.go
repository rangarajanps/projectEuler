package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(FindMaxPalindrome(2))
}

func isPalindrome(num int) bool {

	reverse := 0
	numToReverse := num
	for numToReverse > 0 {
		reverse = reverse*10 + (numToReverse % 10)
		numToReverse = int(numToReverse / 10)
	}

	if num == reverse {
		return true
	}

	return false
}

func FindMaxPalindrome(num int) int {

	if num > 9 {
		fmt.Println("Invalid input greater than 9 provided, exiting the function")
		return -1
	}

	var strIndex string = "9" + strings.Repeat("0", num-1)
	var startIndex, err = strconv.Atoi(strIndex)
	if err != nil {
		fmt.Println("Invalid input passed to the program")
		return -1
	}

	var endStrIndex string = "1" + strings.Repeat("0", num)
	var endIndex, err2 = strconv.Atoi(endStrIndex)
	if err2 != nil {
		fmt.Println("Invalid input passed to the program")
		return -1
	}

	max := 9
	for i := 0; i < num*2-1; i++ {
		max *= 10
	}

	//fmt.Println("max and start, end index -> ", max, startIndex, endIndex)

	for x := startIndex; x < endIndex; x++ {
		for y := startIndex; y < endIndex; y++ {
			temp := x * y
			if isPalindrome(temp) && temp > max {
				max = temp
			}
		}
	}

	return max
}

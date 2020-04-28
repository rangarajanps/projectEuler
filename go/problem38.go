package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(LargestPandigitalMultiples())
}

func LargestPandigitalMultiples() int {

	maxPandigitNum := 1
	for i := 1; i < 20000; i++ {
		temp := ""
		for j := 1; j < 10; j++ {
			prod := i * j
			temp = temp + strconv.Itoa(prod)
			if len(temp) < 9 {
				continue
			} else if len(temp) == 9 {

				if isPandigital(temp) {
					val, err := strconv.Atoi(temp)
					if err != nil {
						fmt.Println("Result will be incorrect, error in converting string to int ", err)
						continue
					}
					if val > maxPandigitNum {
						maxPandigitNum = val
					}
				}

			} else {
				break
			}

		}
	}
	return maxPandigitNum
}

func isPandigital(numAsStr string) bool {

	digits := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := 0; i < len(digits); i++ {
		if strings.Index(numAsStr, digits[i]) == -1 {
			return false
		}
	}
	return true
}

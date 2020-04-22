package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(PandigitProducts())
}

func PandigitProducts() int {

	prodMap := make(map[int]int)
	prod := 1
	sum := 0

	for i := 2; i <= 100; i++ {

		for j := 100; j < 10000; j++ {

			prod = i * j
			pandigitNum := strconv.Itoa(i) + strconv.Itoa(j) + strconv.Itoa(i*j)

			if len(pandigitNum) != 9 {
				continue
			}

			if checkPandigitNumbers(pandigitNum) {
				if _, ok := prodMap[prod]; ok {
					continue
				} else {
					prodMap[prod] = 1
					sum += prod
				}

			}
		}

	}

	return sum
}

func checkPandigitNumbers(number string) bool {

	nums := []rune(number)
	var i rune
	for i = 49; i < 58; i++ { // 49 is 1
		if !contains(nums, i) {
			return false
		}
	}
	return true
  
}

func contains(s []rune, e rune) bool {

	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
  
}

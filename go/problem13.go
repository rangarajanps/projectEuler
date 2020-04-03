package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	test1 := []string{
		`1234`,
		`5678`,
		`9101`,
	}
	fmt.Println(SumOfLargeNumbers(test1...))

	test2 := []string{
		`37107287533902102798797998220837590246510135740250`,
		`46376937677490009712648124896970078050417018260538`,
	}
	fmt.Println(SumOfLargeNumbers(test2...))
  
}

func SumOfLargeNumbers(xi ...string) int {

	strLength := len(xi[0])
	sum := make([]int, strLength+1, strLength+1)
	for i := 0; i < strLength; i++ {
		sum[i] = 0
	}

	for _, v := range xi {
		strArr := strings.Split(v, "")
		temp := 0
		carryFwd := 0
		for itr := strLength - 1; itr >= 0; itr-- {
			value, err := strconv.Atoi(strArr[itr])
			if err != nil {
				fmt.Println("Error in string to int conversion ", err)
			}
			temp = sum[itr+1] + value + carryFwd
			if temp > 9 {
				sum[itr+1] = temp % 10
				carryFwd = temp / 10
			} else {
				sum[itr+1] = temp
				carryFwd = 0
			}
		}
		sum[0] += carryFwd
	}

	//fmt.Println("Sum in string - all digits ", sum)

	// Converting slice of int to slice of string for required 10 or less digits
	limit := 10
	if strLength < 10 {
		limit = strLength + 1
	}

	var total []string
	if sum[0] == 0 {
		limit++
	} else if sum[0] > 10 {
		limit--
	}
	for i := 0; i < limit; i++ {
		if i == 0 && sum[i] == 0 {
			continue
		}
		total = append(total, strconv.Itoa(sum[i]))
	}

	// Converting slice of string to one string
	res := strings.Join(total, "")

	// Converting string to int and returning the value
	result, err := strconv.Atoi(res)
	if err != nil {
		fmt.Printf("String to int conversion error for %v : %v \n", total, err)
		return 0
	}
	return result
}

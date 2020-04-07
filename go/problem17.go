package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(NumberLetterCount(150))
}

func NumberLetterCount(limit int) int {

	sum := 0
	for i := 1; i <= limit; i++ {
		word := convertLetterToWord(i)
		word = strings.ReplaceAll(word, " ", "")
		sum += len(word)
	}
	return sum
}

func convertLetterToWord(n int) string {

	ones := []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
		"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tens := []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

	if n < 20 {
		return ones[n]
	} else if n < 100 {
		tenth := n / 10
		return tens[tenth] + " " + ones[n%10]
	} else if n < 1000 {
		hundreth := n / 100
		if n%100 == 0 {
			return ones[hundreth] + " hundred "
		}
		return ones[hundreth] + " hundred and " + convertLetterToWord(n%100)
	} else if n < 10000 {
		thousanth := n / 1000
		if n%1000 == 0 {
			return ones[thousanth] + "thousand "
		}
		return ones[thousanth] + "thousand and " + convertLetterToWord(n%1000)
	} else {
		million := n / 1000000
		if n%1000000 == 0 {
			return ones[million] + "million "
		}
		return ones[million] + "thousand and " + convertLetterToWord(n%1000000)
	}
}

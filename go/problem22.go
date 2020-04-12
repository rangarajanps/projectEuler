package main

import (
	"fmt"
	"sort"
)

func main() {
	test1 := []string{`THIS`, `IS`, `ONLY`, `A`, `TEST`}
	fmt.Println(NameScores(test1))
}

func NameScores(input []string) int {

	sort.Strings(input)
	sum := 0
	for index, value := range input {
		scr := calculateScore(value)
		sum += (index + 1) * scr
	}
	return sum
}

func calculateScore(word string) int {

	score := 0
	chars := []rune(word)
	for i := 0; i < len(chars); i++ {
		score += int(chars[i]) - 64
	}
	return score
}

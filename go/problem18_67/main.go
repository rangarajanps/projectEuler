package main

import (
	"fmt"
)

func main() {

	inp := [][]int{
		[]int{3, 0, 0, 0},
		[]int{7, 4, 0, 0},
		[]int{2, 4, 6, 0},
		[]int{8, 5, 9, 3},
	}
	fmt.Println(inp)
	fmt.Println(FindMaximumPathSum(inp))
}

func FindMaximumPathSum(intSl [][]int) int {

	for row := len(intSl) - 2; row >= 0; row-- {

		for col := 0; col < row+1; col++ {

			v1 := intSl[row][col] + intSl[row+1][col]
			v2 := intSl[row][col] + intSl[row+1][col+1]
			if v1 > v2 {
				intSl[row][col] = v1
			} else {
				intSl[row][col] = v2
			}

		}
	}

	// fmt.Println(intSl)
	return intSl[0][0]
}

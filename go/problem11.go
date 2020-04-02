package main

import (
	"fmt"
)

func main() {
	testGrid := [][]int{
		[]int{40, 17, 81, 18, 57},
		[]int{74, 4, 36, 16, 29},
		[]int{36, 42, 69, 73, 45},
		[]int{51, 54, 69, 16, 92},
		[]int{7, 97, 57, 32, 16},
	}

	fmt.Println(FindLargestProduct(testGrid))
}

func FindLargestProduct(grid [][]int) int {

	const CONSECUTIVE = 4
	maxProduct := 1

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid); col++ {
			p1, p2, p3, p4 := 1, 1, 1, 1
			for i := 0; i < CONSECUTIVE; i++ {
				p1 *= get(grid, row, col+i)
				p2 *= get(grid, row+i, col)
				p3 *= get(grid, row+i, col+i)
				p4 *= get(grid, row+i, col-i)
			}
			product := max(p1, p2, p3, p4)
			if product > maxProduct {
				maxProduct = product
			}
		}
	}

	return maxProduct
}

func get(arr [][]int, y int, x int) int {

	if 0 <= y && y < len(arr) && 0 <= x && x < len(arr) {
		return arr[y][x]
	}
	return 0
}

func max(xi ...int) int {
	max := xi[0]
	for _, v := range xi[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println(SpiralDiagonalsSum(5))
}

func printArrayAsMatrix(arr [][]int) {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%4d ", arr[i][j])
		}
		fmt.Println()
	}
}

func SpiralDiagonalsSum(limit int) int {

	grid := make([][]int, limit)
	for i := 0; i < limit; i++ {
		grid[i] = make([]int, limit)
	}

	x := limit / 2
	y := x
	num := 1

	for i := 0; i <= limit/2; i++ {

		// Move Right i positions
		for h := 0; h < 2*i+1; h++ {
			grid[x][y] = num
			y = y + 1
			num++
		}
		if num >= limit*limit {
			break
		}

		// Move Down i positions
		for v := 0; v < 2*i+1; v++ {
			grid[x][y] = num
			x = x + 1
			num++
		}

		// Move Left i positions
		for h := 0; h < 2*i+2; h++ {
			grid[x][y] = num
			y = y - 1
			num++
		}

		// Move Up i positions
		for v := 0; v < 2*i+2; v++ {
			grid[x][y] = num
			x = x - 1
			num++
		}

	}

	// fmt.Println("Final Grid is ")
	// printArrayAsMatrix(grid)

	sum := 0

	for p, q := 0, limit-1; p < limit; p, q = p+1, q-1 {
		if p != q {
			sum += grid[p][p] + grid[limit-q-1][q]
		} else {
			sum += grid[p][p]
		}
	}

	return sum
}

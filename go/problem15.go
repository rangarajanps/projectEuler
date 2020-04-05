package main

import (
	"fmt"
)

func main() {
	fmt.Println(LatticePaths(2))
}

func LatticePaths(gridSize int) int {

	n := gridSize * 2
	k := gridSize
	result := 1

	if n-k == 0 {
		k = n - k
	}

	for x := 0; x < k; x++ {
		result = result * (n - x)
		result = result / (x + 1)
	}

	return result

}

package main

import (
	"fmt"
)

func main() {
	fmt.Println(CoinsSum(50))
}

func CoinsSum(n int) int {
	coinValues := []int{1, 2, 5, 10, 20, 50, 100, 200}
	return (splitTotal(n, coinValues, 0))
}

func splitTotal(total int, coins []int, minPos int) int {
	if total == 0 {
		return 1
	}
	count := 0
	for i := minPos; i < len(coins); i++ {
		if coins[i] > total {
			continue
		}
		count += splitTotal(total-coins[i], coins, i)
	}
	return count
}

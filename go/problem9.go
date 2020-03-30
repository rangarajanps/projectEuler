package main

import (
	"fmt"
	"math"
)

func main() {
	ans, err := FindPythogoreanTriplet(24)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ans)
	}
}

func FindPythogoreanTriplet(limit int) (n int, err error) {
	n = 0
	for a := 1; a < int(limit/2); a++ {
		for b := a + 1; b < int(limit/2); b++ {
			c := math.Sqrt(float64(a*a + b*b))
			if c == math.Trunc(c) {
				n = a + b + int(c)
				if n == limit {
					fmt.Println(a, b, int(c))
					n = a * b * int(c)
					return
				}
			}
		}
	}
	err = fmt.Errorf("No such number exist")
	return
}

package main

import (
	"fmt"
)

var numerator = 1
var counter = 0

func main() {
	fmt.Println(ReciprocalCycles(10))
}

func ReciprocalCycles(n int) int {

	rem := 0
	max_denominator := 0

	for denominator := 2; denominator <= n; denominator++ {

		init_rem := numerator % denominator
		counter++
		numerator := init_rem * 10
		rem = init_rem

		for counter < n {

			numerator = rem * 10
			rem = numerator % denominator
			counter++

			if rem == 0 {
				resetCounterNumerator()
				break
			}

			if rem == init_rem {
				if counter >= n {
					return denominator
				}

				if counter > max_denominator {
					max_denominator = denominator
				}

				resetCounterNumerator()
				break
			}

			if counter >= n {
				//fmt.Println("Recurrence NOT detected for denominator "+denominator+" recurring "+counter);
				resetCounterNumerator()
				break
			}
		}

	}
	return max_denominator
}

func resetCounterNumerator() {
	numerator = 1
	counter = 0
}

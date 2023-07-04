package main

import (
	"fmt"
	"strconv"
)

var fibonacci_values []int64

func fibonacci_on_the_fly(n int64) int64 {
	// if fibonacci_values is of length at least n+1, then fibonacci_values[n] has been filled in
	filled_in := int64(len(fibonacci_values)) > n

	// result memoized
	if filled_in {
		return fibonacci_values[n]
	}

	// result not yet memoized
	result := fibonacci_on_the_fly(n-1) + fibonacci_on_the_fly(n-2)
	fibonacci_values = append(fibonacci_values, result)
	return result

}
func main() {
	// Fill-on-the-fly.
	fibonacci_values = make([]int64, 2)
	fibonacci_values[0] = 0
	fibonacci_values[1] = 1

	for {
		// Get n as a string.
		var n_string string
		fmt.Printf("N: ")
		fmt.Scanln(&n_string)

		// If the n string is blank, break out of the loop.
		if len(n_string) == 0 {
			break
		}

		// Convert to int and calculate the Fibonacci number.
		n, _ := strconv.ParseInt(n_string, 10, 64)

		// Uncomment one of the following.
		fmt.Printf("fibonacci_on_the_fly(%d) = %d\n", n, fibonacci_on_the_fly(n))
	}

	// Print out all memoized values just so we can see them.
	for i := 0; i < len(fibonacci_values); i++ {
		fmt.Printf("%d: %d\n", i, fibonacci_values[i])
	}
}

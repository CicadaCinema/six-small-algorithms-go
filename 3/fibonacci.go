package main

import (
	"fmt"
	"strconv"
)

func fibonacci(n int64) int64 {
	// base case
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	// recursive case
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
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
		fmt.Printf("fibonacci(%d) = %d\n", n, fibonacci(n))
	}
}

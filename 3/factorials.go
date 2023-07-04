package main

import "fmt"

func factorial(n int64) int64 {
	// base case
	if n == 0 {
		return 1
	}

	// recursive case
	return n * factorial(n-1)
}

func main() {
	var n int64
	for n = 0; n <= 21; n++ {
		fmt.Printf("%3d! = %20d\n", n, factorial(n))
	}
	fmt.Println()
}

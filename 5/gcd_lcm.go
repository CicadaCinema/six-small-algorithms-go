package main

import (
	"fmt"
	"strconv"
)

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	if a == b {
		return a
	}

	// ensure both arguments are non-negative
	if a < 0 {
		a *= -1
	}
	if b < 0 {
		b *= -1
	}

	var A, B int
	if a > b {
		A = a
		B = b
	} else {
		A = b
		B = a
	}

	// now we have A>B
	R := A % B

	return gcd(B, R)
}

func lcm(a, b int) int {
	return (a / gcd(a, b)) * b
}

func main() {
	for {
		var a_str, b_str string
		fmt.Printf("A: ")
		fmt.Scanln(&a_str)
		fmt.Printf("B: ")
		fmt.Scanln(&b_str)

		// convert to int
		a, _ := strconv.ParseInt(a_str, 10, 0)
		b, _ := strconv.ParseInt(b_str, 10, 0)
		A := int(a)
		B := int(b)

		if a < 1 || b < 1 {
			break
		}

		fmt.Printf("GCD: %d\n", gcd(A, B))
		fmt.Printf("LCM: %d\n", lcm(A, B))
	}
}

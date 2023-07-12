package main

import (
	"fmt"
	"math"
	"strconv"
)

func fast_exp(num, pow int) int {
	result := 1
	for pow > 0 {
		if pow%2 == 1 {
			result *= num
		}
		pow /= 2
		num *= num
	}
	return result
}

func fast_exp_mod(num, pow, mod int) int {
	result := 1
	for pow > 0 {
		if pow%2 == 1 {
			result = (result * num) % mod
		}
		pow /= 2
		num = (num * num) % mod
	}
	return result
}

func main() {
	for {
		var num_str, pow_str, mod_str string
		fmt.Printf("num: ")
		fmt.Scanln(&num_str)
		fmt.Printf("pow: ")
		fmt.Scanln(&pow_str)
		fmt.Printf("mod: ")
		fmt.Scanln(&mod_str)

		// convert to int
		num_, _ := strconv.ParseInt(num_str, 10, 0)
		pow_, _ := strconv.ParseInt(pow_str, 10, 0)
		mod_, _ := strconv.ParseInt(mod_str, 10, 0)

		num := int(num_)
		pow := int(pow_)
		mod := int(mod_)

		if num < 1 || pow < 1 || mod < 1 {
			break
		}

		fmt.Printf("fast_exp: %d\n", fast_exp(num, pow))
		fmt.Printf("fast_exp_mod: %d\n", fast_exp_mod(num, pow, mod))
		real_result := int(math.Pow(float64(num), float64(pow)))
		fmt.Printf("self check: %t\n", fast_exp(num, pow) == real_result && fast_exp_mod(num, pow, mod) == real_result%mod)
	}
}

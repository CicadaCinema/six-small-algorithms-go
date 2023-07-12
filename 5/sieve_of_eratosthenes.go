package main

import (
	"fmt"
	"time"
)

// Build a sieve of Eratosthenes.
func sieve_of_eratosthenes(max int) []bool {
	result := make([]bool, max)
	for i := range result {
		result[i] = true
	}

	// 1 is not prime
	result[0] = false

	current_prime := 2
	for {
		// remove all the multiples of current_prime
		for remove_num := current_prime * 2; remove_num <= max; remove_num += current_prime {
			result[remove_num-1] = false
		}

		// search for the next prime
		found_new := false
		for potential_next := current_prime + 1; potential_next <= max; potential_next++ {
			if result[potential_next-1] {
				found_new = true
				current_prime = potential_next
				break
			}
		}

		// if we have not found any new primes, we are done
		if !found_new {
			break
		}
	}

	return result
}

func print_sieve(sieve []bool) {
	for i, v := range sieve {
		if v {
			fmt.Printf("%d ", i+1)
		}
	}
	fmt.Println()
}

func sieve_to_primes(sieve []bool) []int {
	result := make([]int, 0)
	for i, v := range sieve {
		if v {
			result = append(result, i+1)
		}
	}
	return result
}

func main() {
	var max int
	fmt.Printf("Max: ")
	fmt.Scan(&max)

	start := time.Now()
	sieve := sieve_of_eratosthenes(max)
	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())

	if max <= 1000 {
		print_sieve(sieve)

		primes := sieve_to_primes(sieve)
		fmt.Println(primes)
	}
}

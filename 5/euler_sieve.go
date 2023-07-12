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

func eulers_sieve(max int) []bool {
	result := make([]bool, max)
	for i := range result {
		result[i] = true
	}

	// 1 is not prime
	result[0] = false

	p := 2
	for {
		// remove all the multiples of current_prime, starting from the larger ones

		// integer division really is floor division, so from this we know that q*current_prime >= max
		q := (max / p) + 1
		if q*p > max {
			q--
		}
		// assert multiply_by * current_prime <= max
		// if q*p > max {
		// 	panic("failed assertion 1")
		// }
		for ; q >= p; q-- {
			// if q is not prime, continue
			if result[q-1] {
				// if !result[q*p-1] {
				// 	panic("failed assertion 2")
				// }
				result[q*p-1] = false
			}
		}

		// search for the next prime
		found_new := false
		for potential_next := p + 1; potential_next <= max; potential_next++ {
			if result[potential_next-1] {
				found_new = true
				p = potential_next
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

	fmt.Println("eratosthenes")

	start := time.Now()
	sieve := sieve_of_eratosthenes(max)
	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())

	if max <= 1000 {
		print_sieve(sieve)

		primes := sieve_to_primes(sieve)
		fmt.Println(primes)
	}

	fmt.Println("euler")

	start = time.Now()
	sieve = eulers_sieve(max)
	elapsed = time.Since(start)
	fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())

	if max <= 1000 {
		print_sieve(sieve)

		primes := sieve_to_primes(sieve)
		fmt.Println(primes)
	}
}

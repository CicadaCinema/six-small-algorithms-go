package main

import (
	"fmt"
	"time"
)

func find_factors(num int) []int {
	// create empty slice of ints
	var factors []int

	for num%2 == 0 {
		// num is still divisible by 2
		factors = append(factors, 2)
		num /= 2
	}

	for factor := 3; factor*factor <= num; factor += 2 {
		// if factor is a composite number, then it itself must have prime factors strictly smaller than it
		// and we have already divided num by them in previous iterations of the loop
		for num%factor == 0 {
			// num is still divisible by factor
			factors = append(factors, factor)
			num /= factor
		}

	}

	// at this point we have found all the primes smaller than or equal to the square root of num
	// there can only be one prime factor strictly greater than this
	if num > 1 {
		factors = append(factors, num)
	}

	return factors
}

func find_factors_sieve(num int) []int {
	// create empty slice of ints
	var factors []int

	for i := 0; primes[i]*primes[i] <= num; i++ {
		factor := primes[i]
		// if factor is a composite number, then it itself must have prime factors strictly smaller than it
		// and we have already divided num by them in previous iterations of the loop
		for num%factor == 0 {
			// num is still divisible by factor
			factors = append(factors, factor)
			num /= factor
		}

	}

	// at this point we have found all the primes smaller than or equal to the square root of num
	// there can only be one prime factor strictly greater than this
	if num > 1 {
		factors = append(factors, num)
	}

	return factors
}

func multiply_slice(slice []int) int {
	result := 1
	for _, v := range slice {
		result *= v
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

func sieve_to_primes(sieve []bool) []int {
	result := make([]int, 0)
	for i, v := range sieve {
		if v {
			result = append(result, i+1)
		}
	}
	return result
}

var primes []int

func main() {
	primes = sieve_to_primes(eulers_sieve(2000000000))
	for {
		var num int
		fmt.Printf("Number to factor: ")
		fmt.Scan(&num)

		if num < 2 {
			break
		}

		// Find the factors the slow way.
		start := time.Now()
		factors := find_factors(num)
		elapsed := time.Since(start)
		fmt.Printf("find_factors:       %f seconds\n", elapsed.Seconds())
		// fmt.Println(multiply_slice(factors))
		fmt.Println(factors)
		fmt.Println()

		// Use the Euler's sieve to find the factors.
		start = time.Now()
		factors = find_factors_sieve(num)
		elapsed = time.Since(start)
		fmt.Printf("find_factors_sieve: %f seconds\n", elapsed.Seconds())
		// fmt.Println(multiply_slice(factors))
		fmt.Println(factors)
		fmt.Println()
	}
}

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Initialize a pseudorandom number generator.
var random = rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed

// Return a pseudo random number in the range [min, max).
func rand_range(min int, max int) int {
	return min + random.Intn(max-min)
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

const num_tests = 20

func is_probably_prime(p int) bool {
	for i := 0; i < num_tests; i++ {
		n := rand_range(1, p)
		// this value is always 1 if p is prime
		result_mod_p := fast_exp_mod(n, p-1, p)
		if result_mod_p != 1 {
			return false
		}
	}
	return true
}

// Probabilistically find a prime number within the range [min, max).
func find_prime(min, max int) int {
	for {
		test_prime := rand_range(min, max)
		if is_probably_prime(test_prime) {
			return test_prime
		}
	}
}

func test_known_values() {
	primes := []int{
		10009, 11113, 11699, 12809, 14149,
		15643, 17107, 17881, 19301, 19793,
	}
	composites := []int{
		10323, 11397, 12212, 13503, 14599,
		16113, 17547, 17549, 18893, 19999,
	}

	fmt.Printf(
		"Probability: %f%s\n",
		100-100*math.Pow(0.5, num_tests),
		"%",
	)

	fmt.Println()
	fmt.Println("Primes:")
	for _, v := range primes {
		result := "Composite"
		if is_probably_prime(v) {
			result = "Prime"
		}

		fmt.Printf("%d  %s\n", v, result)
	}

	fmt.Println()
	fmt.Println("Composites:")
	for _, v := range composites {
		result := "Composite"
		if is_probably_prime(v) {
			result = "Prime"
		}

		fmt.Printf("%d  %s\n", v, result)
	}
}

func main() {
	test_known_values()

	for {
		var num_of_digits int
		fmt.Printf(" # Digits: ")
		fmt.Scan(&num_of_digits)

		min := int(math.Pow10(num_of_digits - 1))
		max := int(math.Pow10(num_of_digits))

		// 1 is not a prime number
		if num_of_digits == 1 {
			min = 2
		}

		fmt.Printf(" Prime: %d\n\n", find_prime(min, max))
	}
}

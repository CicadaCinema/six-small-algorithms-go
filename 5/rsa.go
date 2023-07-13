package main

import (
	"fmt"
	"math/rand"
	"time"
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

// Initialize a pseudorandom number generator.
var random = rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed

// Return a pseudo random number in the range [min, max).
func rand_range(min int, max int) int {
	return min + random.Intn(max-min)
}

// Calculate the totient function λ(n)
// where n = p * q and p and q are prime.
func totient(p, q int) int {
	return lcm(p-1, q-1)
}

// Pick a random exponent e in the range (2, lambda_n)
// such that gcd(e, lambda_n) = 1.
func random_exponent(lambda_n int) int {
	for {
		e := rand_range(3, lambda_n)
		if gcd(e, lambda_n) == 1 {
			return e
		}
	}
}

func inverse_mod(a, n int) int {
	t := 0
	newt := 1
	r := n
	newr := a

	for newr != 0 {
		quotient := r / newr
		t, newt = newt, t-quotient*newt
		r, newr = newr, r-quotient*newr
	}

	if r > 1 {

		panic("a is not invertible")
	}
	if t < 0 {
		t = t + n
	}

	return t
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

func main() {
	rand.Seed(time.Now().UnixNano())

	p := find_prime(10000, 50000)
	q := find_prime(10000, 50000)

	// this is the public key modulus
	n := p * q

	lambda_n := totient(p, q)

	// this is the public key exponent
	e := random_exponent(lambda_n)

	d := inverse_mod(e, lambda_n)

	fmt.Println("*** Public ***")
	fmt.Printf("Public key modulus:    %d\n", n)
	fmt.Printf("Public key exponent e: %d\n", e)

	fmt.Println("*** Private ***")
	fmt.Printf("Primes:    %d, %d\n", p, q)
	fmt.Printf("λ(n):      %d\n", lambda_n)
	fmt.Printf("d:         %d\n", d)
	fmt.Println()

	for {
		var m int
		fmt.Printf("Message:    ")
		fmt.Scan(&m)

		if m < 1 {
			break
		}

		ciphertext := fast_exp_mod(m, e, n)
		fmt.Printf("Ciphertext: %d\n", ciphertext)

		plaintext := fast_exp_mod(ciphertext, d, n)
		fmt.Printf("Plaintext:  %d\n", plaintext)

		fmt.Println()
	}
}

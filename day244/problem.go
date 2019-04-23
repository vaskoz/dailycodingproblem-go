package day244

import "math"

// SieveOfEratosthenes returns all the primes strictly less than n.
// If n happens to be a prime, it will NOT be returned in the result.
// Runtime is O(N log log N)
func SieveOfEratosthenes(n int) []int {
	if n < 2 {
		return nil
	}
	isNotPrime := make([]bool, n)
	isNotPrime[0], isNotPrime[1] = true, true
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if !isNotPrime[i] {
			for j := i * i; j < n; j += i {
				isNotPrime[j] = true
			}
		}
	}
	var primes []int
	for p, notPrime := range isNotPrime {
		if !notPrime {
			primes = append(primes, p)
		}
	}
	return primes
}

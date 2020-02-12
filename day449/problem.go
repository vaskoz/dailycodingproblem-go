package day449

import (
	"math"
)

// Sieve of prime numbers up to n.
// Runs in O(N log log N)
func Sieve(n int) []int {
	data := make([]bool, n+1)

	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if !data[i] {
			for j := i * i; j <= n; j += i {
				data[j] = true
			}
		}
	}

	var results []int

	for i := 2; i <= n; i++ {
		if !data[i] {
			results = append(results, i)
		}
	}

	return results
}

// GoldbachsConjecture returns two prime numbers whose sum
// is equal to the given number.
// If even param is less than 4 or an odd number, nil is returned.
func GoldbachsConjecture(even int) []int {
	if even < 4 || even%2 != 0 {
		return nil
	}

	primes := Sieve(even)
	set := make(map[int]struct{}, len(primes))

	for _, prime := range primes {
		set[prime] = struct{}{}
	}

	result := make([]int, 2)

	for _, prime := range primes {
		if _, found := set[even-prime]; found {
			result[0], result[1] = prime, even-prime
			break
		}
	}

	return result
}

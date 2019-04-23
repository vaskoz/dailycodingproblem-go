package day244

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	n      int
	primes []int
}{
	{100, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},
	{2, nil},
	{1, nil},
	{0, nil},
}

func TestSieveOfEratosthenes(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if primes := SieveOfEratosthenes(tc.n); !reflect.DeepEqual(primes, tc.primes) {
			t.Errorf("Expected %v, got %v", tc.primes, primes)
		}
	}
}

func BenchmarkSieveOfEratosthenes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SieveOfEratosthenes(tc.n)
		}
	}
}

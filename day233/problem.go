package day233

// FastFibonnaci runs in O(N) time with O(1) space.
func FastFibonnaci(n uint) uint {
	if n < 2 {
		return n
	}
	first, second := uint(0), uint(1)
	for i := uint(1); i < n; i++ {
		first, second = second, first+second
	}
	return second
}

package day217

// SmallestSparseNumberGreaterThan returns the smallest
// number with no adjacent 1's in the binary form.
// If n is already sparse, it returns n.
// Runs in O(num of bits)
func SmallestSparseNumberGreaterThan(n uint) uint {
	if n < 3 {
		return n
	}
	var result uint
	var bits, ones int
	for n != 0 {
		bit := n & 1
		n >>= 1
		if bit == 1 {
			ones++
		} else {
			if ones > 1 {
				n <<= 1
				n |= 1
				result = 1 << uint(bits)
				bits--
			} else {
				result |= 1 << uint(bits-1)
			}
			ones = 0
		}
		bits++
	}
	if ones == 1 {
		result |= 1 << uint(bits-1)
	} else if ones > 1 {
		result = 1 << uint(bits)
	}
	return result
}

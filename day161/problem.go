package day161

// Reverse32Bits reverses the bits and returned the result.
// Runs in O(32) time.
func Reverse32Bits(num uint32) uint32 {
	var rev uint32
	offset := uint(32)
	for num != 0 {
		offset--
		rev <<= 1
		if num&1 == 1 {
			rev ^= 1
		}
		num >>= 1
	}
	return rev << offset
}

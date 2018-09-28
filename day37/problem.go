package day37

import "math/bits"

// PowerSet returns all subsets.
func PowerSet(set []int) [][]int {
	result := make([][]int, 0, 1<<uint64(len(set)))
	for i := uint64(0); i < (1 << uint64(len(set))); i++ {
		entry := make([]int, 0, bits.OnesCount64(i))
		for pos := 0; pos < len(set); pos++ {
			if i&(1<<uint64(pos)) != 0 {
				entry = append(entry, set[pos])
			}
		}
		result = append(result, entry)
	}
	return result
}

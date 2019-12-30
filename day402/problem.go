package day402

import (
	"fmt"
	"sort"
)

// StrobogrammaticNumber returns all the strobogrammatic numbers
// with the given number of digits.
func StrobogrammaticNumber(digits int) []string {
	result := strobogrammaticNumber(digits, digits)
	sort.Strings(result) // for deterministic ordering of results

	return result
}

func strobogrammaticNumber(digits, length int) []string {
	if digits == 0 {
		return nil
	} else if digits == 1 {
		return []string{"0", "1", "8"}
	}

	mid := strobogrammaticNumber(digits-2, length)
	if mid == nil {
		mid = []string{""}
	}

	result := make([]string, 0, len(mid)*5)

	for _, m := range mid {
		if digits != length {
			result = append(result, fmt.Sprintf("0%s0", m))
		}

		result = append(result, fmt.Sprintf("1%s1", m))
		result = append(result, fmt.Sprintf("6%s9", m))
		result = append(result, fmt.Sprintf("8%s8", m))
		result = append(result, fmt.Sprintf("9%s6", m))
	}

	return result
}

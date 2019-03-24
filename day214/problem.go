package day214

import "strconv"

// LengthLongestConsecutiveRunOnesBitShift calculates the longest
// run of 1's using bit shifting.
func LengthLongestConsecutiveRunOnesBitShift(n int) int {
	var max, count int
	for n != 0 {
		if n&1 == 1 {
			count++
		} else {
			if count > max {
				max = count
			}
			count = 0
		}
		n >>= 1
	}
	if count > max {
		return count
	}
	return max
}

// LengthLongestConsecutiveRunOnesString converts the number
// into a binary representation string.
func LengthLongestConsecutiveRunOnesString(n int) int {
	var max, count int
	for _, digit := range strconv.FormatInt(int64(n), 2) {
		if digit == '1' {
			count++
		} else {
			if count > max {
				max = count
			}
			count = 0
		}
	}
	if count > max {
		return count
	}
	return max
}

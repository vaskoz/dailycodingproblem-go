package day118

import (
	"sort"
)

// SquareElementsSortedBrute returns the squared result of every element of the param.
// Runs in O(N log N) due to sorting the results.
func SquareElementsSortedBrute(sorted []int) []int {
	result := make([]int, len(sorted))
	for i := range sorted {
		result[i] = sorted[i] * sorted[i]
	}
	sort.Ints(result)
	return result
}

// SquareElementsSortedLinear returns the squared result of every element of the param.
// Runs in O(N) due to merging two sorted lists into the result.
func SquareElementsSortedLinear(sorted []int) []int {
	result := make([]int, 0, len(sorted))
	firstPositiveIndex := 0
	for i := range sorted {
		if sorted[i] >= 0 {
			firstPositiveIndex = i
			break
		}
	}
	negI, posI := firstPositiveIndex-1, firstPositiveIndex
	for {
		if negI >= 0 && posI < len(sorted) {
			a := sorted[negI] * sorted[negI]
			b := sorted[posI] * sorted[posI]
			if a < b {
				result = append(result, a)
				negI--
			} else {
				result = append(result, b)
				posI++
			}
		} else if negI >= 0 {
			result = append(result, sorted[negI]*sorted[negI])
			negI--
		} else if posI < len(sorted) {
			result = append(result, sorted[posI]*sorted[posI])
			posI++
		} else {
			break
		}
	}
	return result
}

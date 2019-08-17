package day359

import "fmt"

// ErrNotPossible returns the error associated with an
// impossible input.
func ErrNotPossible() error {
	return errNotPossible
}

var errNotPossible = fmt.Errorf("no solution possible")

// ConvertAnagramToInt takes a concatenation
// of integers and returns a sorted list of decimal integers.
// Panics if impossible due to too many letters or incorrect
// letters.
func ConvertAnagramToInt(str string) (string, error) {
	freq := make(map[rune]int)
	for _, r := range str {
		freq[r]++
	}
	nums := []string{"zero", "one", "two", "three", "four",
		"five", "six", "seven", "eight", "nine"}
	numFreq := make([]map[rune]int, len(nums))
	for i := range nums {
		numFreq[i] = make(map[rune]int)
		for _, r := range nums[i] {
			numFreq[i][r]++
		}
	}
	return convertAnagramToInt(freq, numFreq, "")
}

// Runs DFS until a solution is found.
func convertAnagramToInt(freq map[rune]int, nums []map[rune]int, result string) (string, error) {
	baseCase := true
	for _, c := range freq {
		if c != 0 {
			baseCase = false
			break
		}
	}
	if baseCase {
		return result, nil
	}
	var digit int
	for digit < 10 {
		valid := true
		for r, c := range nums[digit] {
			if freq[r]-c < 0 {
				valid = false
				break
			}
		}
		if valid {
			for r, c := range nums[digit] {
				freq[r] -= c
			}
			r, err := convertAnagramToInt(freq, nums, fmt.Sprintf("%s%d", result, digit))
			if err == nil {
				return r, nil
			}
			for r, c := range nums[digit] {
				freq[r] += c
			}
		}
		digit++
	}
	return "", errNotPossible
}

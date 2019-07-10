package day322

// NumberOfJumps returns the number of jumps to reach
// N from 0.
// Each jump length is of length "i".
// Runs in O(answer) time and O(1) space.
func NumberOfJumps(n int) int {
	if n < 0 {
		n = -n
	}
	var result int
	for (arithmeticSum(result) < n) || ((arithmeticSum(result)-n)&1 == 1) {
		result++
	}
	return result
}

// arithmeticSum is:
// https://en.wikipedia.org/wiki/Arithmetic_progression#Sum
func arithmeticSum(n int) int {
	return (n * (n + 1)) / 2
}

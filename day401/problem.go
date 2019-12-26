package day401

// ApplyPermutation mutates the given input based on the
// permutation information in the 2nd argument.
// Runs in O(N) time and O(N/2) space.
// Correct input is the responsibility of the caller.
func ApplyPermutation(input []interface{}, permutation []int) {
	if len(input) != len(permutation) {
		panic("permutation length must match input length")
	}

	passed := make(map[int]interface{}, len(input))

	for i, pos := range permutation {
		if pos < i {
			input[i] = passed[pos]
		} else {
			passed[i] = input[i]
			input[i] = input[pos]
		}
	}
}

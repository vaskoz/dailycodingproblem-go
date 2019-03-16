package day206

// ReorderPermutation reorders the input based on the positions
// provided in the 'permutationPosition' slice.
// Runs in O(N) and O(N) space for result.
func ReorderPermutation(input []interface{}, permutationPosition []int) []interface{} {
	result := make([]interface{}, len(input))
	for to, from := range permutationPosition {
		result[to] = input[from]
	}
	return result
}

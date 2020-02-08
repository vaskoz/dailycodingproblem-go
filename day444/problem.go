package day444

// StringMatchKMP is an implementation of
// https://en.wikipedia.org/wiki/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm
// NOTE: only returns the index of the first occurrence. If not found, returns -1.
// WARNING: only works on ascii characters.
func StringMatchKMP(input, substr string) int {
	var i, j int

	table := buildTable(substr)

	for i < len(input) {
		if substr[j] == input[i] {
			i++
			j++
		}

		if j == len(substr) {
			return i - j
		}

		if i < len(input) && substr[j] != input[i] {
			if j != 0 {
				j = table[j-1]
			} else {
				i++
			}
		}
	}

	return -1
}

func buildTable(substr string) []int {
	table := make([]int, len(substr))
	i := 1
	last := 0

	for i < len(table) {
		if substr[i] == substr[last] {
			last++
			table[i] = last
			i++
		} else {
			if last != 0 {
				last = table[last-1]
			} else {
				table[i] = 0
				i++
			}
		}
	}

	return table
}

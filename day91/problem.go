package day91

// BadSnippet is given in the problem.
func BadSnippet() []int {
	funcs := make([]func() int, 9)
	for i := 0; i < 9; i++ {
		funcs[i] = func() int {
			return i //nolint: scopelint
		}
	}
	result := make([]int, 9)
	for i, f := range funcs {
		result[i] = f()
	}
	return result
}

// FixedSnippet is how it should be written.
func FixedSnippet() []int {
	funcs := make([]func(int) int, 9)
	for i := 0; i < 9; i++ {
		funcs[i] = func(i int) int {
			return i
		}
	}
	result := make([]int, 9)
	for i, f := range funcs {
		result[i] = f(i)
	}
	return result
}

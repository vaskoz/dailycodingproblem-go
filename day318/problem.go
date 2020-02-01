package day318

// NumberOfValidPlayListsRec is a recursive solution.
func NumberOfValidPlayListsRec(n, m, b int) int {
	if b >= m {
		return 0
	}

	if n == 0 {
		return 1
	}

	mp := m

	if b != 0 {
		mp--
	}

	return m * NumberOfValidPlayListsRec(n-1, mp, max(0, b-1))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// NumberOfValidPlayListsIterative is an iterative solution.
func NumberOfValidPlayListsIterative(n, m, b int) int {
	if b >= m {
		return 0
	}

	total := 1
	available := m

	for i := 0; i <= b; i++ {
		total *= available
		available--
	}

	if available < m {
		available++
	}

	for i := b + 1; i < n; i++ {
		total *= available
	}

	return total
}

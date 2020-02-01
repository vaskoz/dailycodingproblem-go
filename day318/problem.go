package day318

// NumberOfValidPlayListsRec is a recursive solution.
func NumberOfValidPlayListsRec(required, downloaded, buffer int) int {
	if buffer >= downloaded {
		return 0
	}

	if required == 0 {
		return 1
	}

	downloadParam := downloaded

	if buffer != 0 {
		downloadParam--
	}

	return downloaded * NumberOfValidPlayListsRec(required-1, downloadParam, max(0, buffer-1))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// NumberOfValidPlayListsIterative is an iterative solution.
func NumberOfValidPlayListsIterative(required, downloaded, buffer int) int {
	if buffer >= downloaded {
		return 0
	}

	total := 1
	available := downloaded

	for i := 0; i <= buffer; i++ {
		total *= available
		available--
	}

	if available < downloaded {
		available++
	}

	for i := buffer + 1; i < required; i++ {
		total *= available
	}

	return total
}

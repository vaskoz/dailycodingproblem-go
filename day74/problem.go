package day74

// MultiplicationTableBrute counts the number of X occurrences
// in a N by N multiplication table.
// Runs in O(N^2) time.
func MultiplicationTableBrute(N, X int) int {
	count := 0
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			if i*j == X {
				count++
			}
		}
	}
	return count
}

// MultiplicationTableLinear counts the number of X occurrences
// in a N by N multiplication table.
// Runs in O(N) time.
func MultiplicationTableLinear(N, X int) int {
	count := 0
	for i := 1; i <= N; i++ {
		if X%i == 0 && X/i <= N {
			count++
		}
	}
	return count
}

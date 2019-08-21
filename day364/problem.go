package day364

// LongestIncreasingSubsequence returns the longest increasing
// subsequence.
// Runs in O(N log N) time with O(N) extra space.
func LongestIncreasingSubsequence(seq []int) []int {
	p := make([]int, len(seq))
	m := make([]int, len(seq)+1)
	var l int
	for i := range seq {
		lo := 1
		hi := l
		for lo <= hi {
			mid := (lo + hi) / 2
			if (lo+hi)%2 != 0 {
				mid++
			}
			if seq[m[mid]] <= seq[i] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		newL := lo
		p[i] = m[newL-1]
		m[newL] = i
		if newL > l {
			l = newL
		}
	}
	s := make([]int, l)
	k := m[l]
	for i := len(s) - 1; i >= 0; i-- {
		s[i] = seq[k]
		k = p[k]
	}
	return s
}

package day75

// LongestIncreasingSubsequence returns the longest increasing subsequence.
// Runs in O(N log N) time.
func LongestIncreasingSubsequence(x []int) []int {
	p := make([]int, len(x))
	m := make([]int, len(x)+1)
	l := 0
	for i := range x {
		lo := 1
		hi := l
		for lo <= hi {
			var mid int
			if (lo+hi)%2 == 0 {
				mid = (lo + hi) / 2
			} else {
				mid = ((lo + hi) / 2) + 1
			}
			if x[m[mid]] < x[i] {
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
	for i := range s {
		s[len(s)-1-i] = x[k]
		k = p[k]
	}
	return s
}

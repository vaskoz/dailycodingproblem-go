package day264

// DeBruijnSequence returns a cyclic sequence in
// which every possible k-length string of characters
// in C occurs exactly once.
func DeBruijnSequence(c string, k int) string {
	alphabet := []rune(c)
	a := make([]int, len(alphabet)*k)
	var seq []int
	var db func(int, int)
	db = func(t, p int) {
		if t > k {
			if k%p == 0 {
				seq = append(seq, a[1:p+1]...)
			}
		} else {
			a[t] = a[t-p]
			db(t+1, p)
			for j := a[t-p] + 1; j < len(c); j++ {
				a[t] = j
				db(t+1, t)
			}
		}
	}
	db(1, 1)
	result := make([]rune, len(seq))
	for i, v := range seq {
		result[i] = alphabet[v]
	}
	return string(result)
}

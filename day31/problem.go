package day31

// EditDistance returns the minimum number of additions, removals,
// and substitutions required to make the two strings equal.
// Runtime O(len(s)*len(t)) and space O(len(s)*len(t))
func EditDistance(s, t string) int {
	d := make([][]int, len(s)+1)
	for i := range d {
		d[i] = make([]int, len(t)+1)
		d[i][0] = i
	}
	for j := range d[0] {
		d[0][j] = j
	}
	for j, tr := range t {
		for i, sr := range s {
			if sr == tr {
				d[i+1][j+1] = d[i][j]
			} else {
				d[i+1][j+1] = minimum(d[i][j+1]+1,
					d[i+1][j]+1,
					d[i][j]+1)
			}
		}
	}
	return d[len(s)][len(t)]
}

func minimum(a, b, c int) int {
	result := a
	if b < result {
		result = b
	}
	if c < result {
		result = c
	}
	return result
}

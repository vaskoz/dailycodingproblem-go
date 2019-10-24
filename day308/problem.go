package day308

// CountParensEqualTrue returns the number of possible
// paren patterns that evaluate to True.
// Runs in O(N^3) time and O(N^2) extra space.
func CountParensEqualTrue(input []rune) int {
	var sym, ops []rune

	for i := range input {
		if i%2 == 0 {
			sym = append(sym, input[i])
		} else {
			ops = append(ops, input[i])
		}
	}

	f := make([][]int, len(sym))
	t := make([][]int, len(sym))

	for i := range f {
		f[i] = make([]int, len(sym))
		t[i] = make([]int, len(sym))

		if sym[i] == 'F' {
			f[i][i] = 1
		}

		if sym[i] == 'T' {
			t[i][i] = 1
		}
	}

	for gap := 1; gap < len(sym); gap++ {
		for i, j := 0, gap; j < len(sym); i, j = i+1, j+1 {
			t[i][j] = 0
			f[i][j] = 0

			for g := 0; g < gap; g++ {
				k := i + g
				tik := t[i][k] + f[i][k]
				tkj := t[k+1][j] + f[k+1][j]

				if ops[k] == '&' {
					t[i][j] += t[i][k] * t[k+1][j]
					f[i][j] += tik*tkj - t[i][k]*t[k+1][j]
				}

				if ops[k] == '|' {
					f[i][j] += f[i][k] * f[k+1][j]
					t[i][j] += tik*tkj - f[i][k]*f[k+1][j]
				}

				if ops[k] == '^' {
					t[i][j] += f[i][k]*t[k+1][j] + t[i][k]*f[k+1][j]
					f[i][j] += t[i][k]*t[k+1][j] + f[i][k]*f[k+1][j]
				}
			}
		}
	}

	return t[0][len(t[0])-1]
}

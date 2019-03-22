package day212

// SpreadsheetColumnName returns the name of the column
// based on the ordinal position.
func SpreadsheetColumnName(pos int) string {
	if pos < 1 {
		return ""
	}
	pos--
	var ans []rune
	for pos >= 0 {
		ans = append(ans, 'A'+rune(pos%26))
		pos /= 26
		if pos == 0 {
			pos = -1
		}
		pos--
	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return string(ans)
}

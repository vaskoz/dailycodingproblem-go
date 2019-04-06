package day142

// WildcardParens answers if the parens are balanced.
// This includes a wildcard character '*' that can be either paren
// or an empty string.
func WildcardParens(input string) bool {
	var opens, wildcards int
	for _, c := range input {
		switch c {
		case '(':
			opens++
		case ')':
			switch {
			case opens == 0 && wildcards < 1:
				return false
			case opens == 0:
				wildcards--
			default:
				opens--
			}
		case '*':
			wildcards++
		}
	}
	return opens == 0 || opens <= wildcards
}

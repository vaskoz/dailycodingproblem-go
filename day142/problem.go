package day142

// WildcardParens answers if the parens are balanced.
// This includes a wildcard character '*' that can be either paren
// or an empty string.
func WildcardParens(input string) bool {
	var opens, wildcards int
	for _, c := range input {
		if c == '(' {
			opens++
		} else if c == ')' {
			if opens == 0 && wildcards < 1 {
				return false
			} else if opens == 0 {
				wildcards--
			} else {
				opens--
			}
		} else if c == '*' {
			wildcards++
		}
	}
	return opens == 0 || opens <= wildcards
}

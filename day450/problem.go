package day450

// WildcardParens answers if the parens are balanced.
// This includes a wildcard character '*' that can be either paren
// or an empty string.
// Implementation note: counts opens two different ways.
// First, it counts opens assuming a wildcard is an open '('.
// Next, it counts opens assuming a wildcard is a close ')'.
// Runs in O(N) time and O(1) space.
func WildcardParens(input string) bool {
	var opensIfWildcardCloses, opensIfWildcardOpens int

	for _, c := range input {
		switch c {
		case '(':
			opensIfWildcardCloses++
			opensIfWildcardOpens++
		case ')':
			opensIfWildcardCloses--
			opensIfWildcardOpens--
		case '*':
			opensIfWildcardCloses--
			opensIfWildcardOpens++
		}

		if opensIfWildcardOpens < 0 {
			break
		}
		// if a closing wildcard goes negative,
		// assume it's an empty string.
		if opensIfWildcardCloses < 0 {
			opensIfWildcardCloses = 0
		}
	}

	return opensIfWildcardCloses == 0
}

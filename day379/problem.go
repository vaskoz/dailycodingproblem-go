package day379

// AllPossibleSubsequences returns all the possible
// ordered subsequences.
func AllPossibleSubsequences(str string) map[string]struct{} {
	ansSet := make(map[string]struct{})
	allPossibleSubsequences(str, ansSet)

	return ansSet
}

func allPossibleSubsequences(str string, ansSet map[string]struct{}) {
	if len(str) == 0 {
		return
	}

	for i := range str {
		sub := str[:i] + str[i+1:]
		allPossibleSubsequences(sub, ansSet)
	}

	ansSet[str] = struct{}{}
}

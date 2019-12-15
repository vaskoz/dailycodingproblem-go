package day391

// LongestContiguousBrowsingHistory returns the longest contiguous
// common search history.
// Runs in O(N*k) time where K is the length of the answer.
func LongestContiguousBrowsingHistory(user1, user2 []string) []string {
	first := make(map[string]struct{})
	mutual := make(map[string]struct{})

	for _, path := range user1 {
		first[path] = struct{}{}
	}

	for _, path := range user2 {
		if _, found := first[path]; found {
			mutual[path] = struct{}{}
		}
	}

	var longest []string

	for i, path1 := range user1 {
		if _, common := mutual[path1]; !common {
			continue
		}

		for j, path2 := range user2 {
			if path2 == path1 {
				ptrI := i
				ptrJ := j

				for ptrI < len(user1) && ptrJ < len(user2) && user1[ptrI] == user2[ptrJ] {
					ptrI++
					ptrJ++
				}

				if ptrI-i > len(longest) {
					longest = user1[i:ptrI]
				}
			}
		}
	}

	return longest
}

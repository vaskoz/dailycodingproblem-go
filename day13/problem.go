package day13

// LongestSubKDistinct returns  find the length of the longest substring
// that contains at most k distinct characters.
// Runs in O(N) time. Uses O(K) space for the map.
func LongestSubKDistinct(s string, k int) string {
	r := []rune(s)
	if len(r) <= k {
		return string(r)
	}
	letters := make(map[rune]int)
	var curL, maxL, curR, maxR int
	for curR < len(r) {
		if _, found := letters[r[curR]]; len(letters) < k || found {
			letters[r[curR]]++
			curR++
			if curR-curL > maxR-maxL {
				maxL, maxR = curL, curR
			}
		} else {
			for len(letters) >= k {
				if count := letters[r[curL]]; count == 1 {
					delete(letters, r[curL])
				} else {
					letters[r[curL]]--
				}
				curL++
			}
		}

	}
	return string(r[maxL:maxR])
}

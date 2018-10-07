package day46

// IsPalindrome answers this question is linear time O(N) and O(1) space.
func IsPalindrome(letters []rune) bool {
	for i := 0; i < len(letters)/2; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

// LongestPalindromicSubstringBrute runs in O(N^3) time and O(1) space.
func LongestPalindromicSubstringBrute(str string) string {
	letters := []rune(str)
	for l := len(letters); l > 0; l-- {
		for left, right := 0, l; right <= len(letters); left, right = left+1, right+1 {
			if IsPalindrome(letters[left:right]) {
				return string(letters[left:right])
			}
		}
	}
	return ""
}

// LongestPalindromicSubstring runs in O(N) time and O(N) space.
func LongestPalindromicSubstring(str string) string {
	letters := []rune(str)
	s2 := addBoundaries(letters)
	p := make([]int, len(s2))

	var c, r, m, n int
	for i := 1; i < len(s2); i++ {
		if i > r {
			p[i] = 0
			m = i - 1
			n = i + 1
		} else {
			i2 := c*2 - i
			if p[i2] < r-i-1 {
				p[i] = p[i2]
				m = -1
			} else {
				p[i] = r - i
				n = r + 1
				m = i*2 - n
			}
		}
		for m >= 0 && n < len(s2) && s2[m] == s2[n] {
			p[i]++
			m--
			n++
		}
		if i+p[i] > r {
			c = i
			r = i + p[i]
		}
	}
	c = 0
	l := 0
	for i := 1; i < len(s2); i++ {
		if l < p[i] {
			l = p[i]
			c = i
		}
	}
	return string(removeBoundaries(s2[c-l : c+l+1]))
}

func addBoundaries(letters []rune) []rune {
	if len(letters) == 0 {
		return []rune{'|', '|'}
	}
	cs := make([]rune, len(letters)*2+1)
	for i := 0; i < len(cs)-1; i += 2 {
		cs[i] = '|'
		cs[i+1] = letters[i/2]
	}
	cs[len(cs)-1] = '|'
	return cs
}

func removeBoundaries(letters []rune) []rune {
	if len(letters) < 3 {
		return []rune{}
	}
	cs := make([]rune, (len(letters)-1)/2)
	for i := range cs {
		cs[i] = letters[i*2+1]
	}
	return cs
}

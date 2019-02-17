package day176

// DoesMappingExist returns true if a 1-to-1 mapping exists.
func DoesMappingExist(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m := make(map[rune]rune)
	for i, r := range s1 {
		if t, exists := m[r]; exists && rune(s2[i]) != t {
			return false
		} else if !exists {
			m[r] = rune(s2[i])
		}
	}
	return true
}

package day250

import "sort"

// CryptarithmeticAdditionPuzzle returns a possible solution to the addition.
func CryptarithmeticAdditionPuzzle(first, second, result string) map[rune]int {
	letters := make(map[rune]struct{}, 10)
	for _, part := range []string{first, second, result} {
		for _, r := range part {
			letters[r] = struct{}{}
		}
	}
	if len(letters) > 10 {
		return nil
	}
	runes := make([]rune, 0, len(letters))
	for r := range letters {
		runes = append(runes, r)
	}
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	availDigits := make([]bool, 10)
	for i := 0; i < 10; i++ {
		availDigits[i] = true
	}
	return backtrackingSearch(runes, make(map[rune]int), availDigits, first, second, result)
}

func backtrackingSearch(runes []rune, visited map[rune]int,
	availDigits []bool, first, second, result string) map[rune]int {
	if len(runes) == 0 && validSolution(first, second, result, visited) {
		return visited
	} else if len(runes) == 0 {
		return nil
	}
	nextRune := runes[0]
	for d, avail := range availDigits {
		if !avail {
			continue
		}
		availDigits[d] = false
		visited[nextRune] = d
		if result := backtrackingSearch(runes[1:], visited, availDigits, first, second, result); result != nil {
			return result
		}
		delete(visited, nextRune)
		availDigits[d] = true
	}
	return nil
}

func validSolution(first, second, result string, visited map[rune]int) bool {
	f := []rune(first)
	s := []rune(second)
	t := []rune(result)
	exp := 1
	var fi int
	for i := len(f) - 1; i >= 0; i-- {
		fi += visited[f[i]] * exp
		exp *= 10
	}
	exp = 1
	var si int
	for i := len(s) - 1; i >= 0; i-- {
		si += visited[s[i]] * exp
		exp *= 10
	}
	exp = 1
	var ti int
	for i := len(t) - 1; i >= 0; i-- {
		ti += visited[t[i]] * exp
		exp *= 10
	}
	return fi+si == ti
}

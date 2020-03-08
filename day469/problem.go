package day469

import "fmt"

// Guess represents a code and a score for that guess.
type Guess struct {
	Code  string
	Score int
}

// ValidMastermindGuesses returns true if a code exists that
// correspond to a valid Mastermind code.
// Runs in O(151,200*N) time.
// There are 151,200 permutations of 6 unique digits.
// 10*9*8*7*6*5 == 151,200 permutations. N is the number of
// guesses.
func ValidMastermindGuesses(guesses []Guess) bool {
	for _, p := range perm6([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		valid := true

		for _, g := range guesses {
			var count int

			for i := 0; i < len(p); i++ {
				if p[i] == g.Code[i] {
					count++
				}
			}

			if count != g.Score {
				valid = false
				break
			}
		}

		if valid {
			return true
		}
	}

	return false
}

// ValidMastermindGuessesAll returns all the codes that
// correspond to a valid Mastermind code.
// Runs in O(151,200*N) time.
// There are 151,200 permutations of 6 unique digits.
// 10*9*8*7*6*5 == 151,200 permutations. N is the number of
// guesses.
func ValidMastermindGuessesAll(guesses []Guess) []string {
	var result []string

	for _, p := range perm6([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		valid := true

		for _, g := range guesses {
			count := 0

			for i := 0; i < len(p); i++ {
				if p[i] == g.Code[i] {
					count++
				}
			}

			if count != g.Score {
				valid = false
				break
			}
		}

		if valid {
			result = append(result, p)
		}
	}

	return result
}

func perm6(digits []int) []string {
	if len(digits) == 4 {
		return []string{""}
	}

	var result []string

	for i, d := range digits {
		ds := fmt.Sprintf("%d", d)

		var subdigits []int
		subdigits = append(subdigits, digits[:i]...)
		subdigits = append(subdigits, digits[i+1:]...)

		for _, p := range perm6(subdigits) {
			result = append(result, ds+p)
		}
	}

	return result
}

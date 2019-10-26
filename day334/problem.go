package day334

const goal = 24

// nolint
var operators = []rune{'+', '-', '*', '/'}

// TwentyFourGame returns true if the game of 24
// is possible given this input.
// Operators used are: +, -, *, and /
func TwentyFourGame(digits []int) bool {
	if len(digits) == 1 {
		return digits[0] == goal
	}

	for i := 0; i < len(digits)-1; i++ {
		for _, op := range operators {
			var next []int
			next = append(next, digits[:i]...)

			switch op {
			case '+':
				next = append(next, digits[i]+digits[i+1])
			case '-':
				next = append(next, digits[i]-digits[i+1])
			case '*':
				next = append(next, digits[i]*digits[i+1])
			case '/':
				if digits[i+1] == 0 {
					continue
				}

				next = append(next, digits[i]/digits[i+1])
			}

			next = append(next, digits[i+2:]...)

			if TwentyFourGame(next) {
				return true
			}
		}
	}

	return false
}

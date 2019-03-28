package day216

var romanNumeralMap = map[rune]int{
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	'V': 5,
	'I': 1,
}

// ConvertRomanNumeralToInteger converts and roman numeral string
// into an integer.
// Only works on non-negative integers (no decimals).
func ConvertRomanNumeralToInteger(roman string) int {
	var result int
	for i := 0; i < len(roman); i++ {
		if first := romanNumeralMap[rune(roman[i])]; i+1 < len(roman) {
			if second := romanNumeralMap[rune(roman[i+1])]; first >= second {
				result += first
			} else {
				result += second - first
				i++
			}
		} else {
			result += first
			i++
		}
	}
	return result
}

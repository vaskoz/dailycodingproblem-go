package day81

// DigitMappings represents digits that can be mapped to multiple other values.
// Think of the alpha codes on a phone pad.
type DigitMappings map[rune][]string

// PossiblePhoneLetters takes a given phone number and returns all possible letter
// combinations.
func PossiblePhoneLetters(number string, digitMappings DigitMappings) []string {
	var result []string
	for i, r := range number {
		letters := digitMappings[r]
		if i == 0 {
			result = append(result, letters...)
		} else {
			tmp := make([]string, 0, len(letters)*len(result))
			for _, part := range result {
				for _, letter := range letters {
					tmp = append(tmp, part+letter)
				}
			}
			result = tmp
		}
	}
	return result
}

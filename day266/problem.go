package day266

import "reflect"

// AllValidStepWords returns all valid step words by anagramming
// an appended letter.
// Runs in O(dict*26)
func AllValidStepWords(dict []string, input string) []string {
	dictAnagramData := make([]map[rune]int, len(dict))
	for i, word := range dict {
		dictAnagramData[i] = make(map[rune]int)
		for _, r := range word {
			dictAnagramData[i][r]++
		}
	}
	inputAnagramData := make(map[rune]int)
	for _, r := range input {
		inputAnagramData[r]++
	}
	var result []string
	for r := 'a'; r <= 'z'; r++ {
		inputAnagramData[r]++
		for i, data := range dictAnagramData {
			if reflect.DeepEqual(data, inputAnagramData) {
				result = append(result, dict[i])
			}
		}
		inputAnagramData[r]--
		if inputAnagramData[r] == 0 {
			delete(inputAnagramData, r)
		}
	}
	return result
}

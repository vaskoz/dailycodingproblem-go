package day22

import "testing"

var testcases = []struct {
	words        WordSet
	concatenated string
	expected     string
	expectedErr  error
}{
	{WordSet{"quick": struct{}{}, "brown": struct{}{}, "fox": struct{}{}, "the": struct{}{}},
		"thequickbrownfox", "the quick brown fox", nil},
	{WordSet{"quick": struct{}{}, "brown": struct{}{}, "fox": struct{}{}, "the": struct{}{}},
		"thequickbrownfoxy", "", errNoSentencePossible},
}

func TestOriginalSentence(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result, err := OriginalSentence(tc.words, tc.concatenated); result != tc.expected || err != tc.expectedErr {
			t.Errorf("Expected (%v, %v) but got (%v, %v)", tc.expected, tc.expectedErr, result, err)
		}
	}
}

func BenchmarkOriginalSentence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			OriginalSentence(tc.words, tc.concatenated)
		}
	}
}

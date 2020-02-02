package day431

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	text      string
	sentences []string
}{
	{
		"this is not valid. However, this sentence is totally valid. Is it not?",
		[]string{
			"However, this sentence is totally valid.",
			"Is it not?",
		},
	},
	{
		"Numbers like 9 and extra  spaces. A bee bit me!",
		[]string{
			"A bee bit me!",
		},
	},
	{
		"A bee!",
		[]string{
			"A bee!",
		},
	},
	{
		"The colors: red, white and blue. I can not : space punctuations.",
		[]string{
			"The colors: red, white and blue.",
		},
	},
	{
		"I do not allow CamelCase. It is not good for anybody, right?",
		[]string{
			"It is not good for anybody, right?",
		},
	},
}

func TestValidSentences(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := ValidSentences(tc.text); !reflect.DeepEqual(result, tc.sentences) {
			t.Errorf("Expected %v, got %v", tc.sentences, result)
		}
	}
}

func BenchmarkValidSentences(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ValidSentences(tc.text)
		}
	}
}

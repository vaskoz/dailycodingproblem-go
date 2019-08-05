package day345

import "testing"

// nolint
var testcases = []struct {
	sentence1, sentence2 string
	thesaurus            Thesaurus
	expected             bool
}{
	{
		"he wants to eat food",
		"he wants to consume food",
		Thesaurus{
			"big": {
				"large": struct{}{},
			},
			"eat": {
				"consume": struct{}{},
			},
		},
		true,
	},
	{
		"he wants to eat food",
		"he wants to devour food",
		Thesaurus{
			"eat": {
				"consume": struct{}{},
			},
		},
		false,
	},
	{
		"he wants to eat food",
		"he wants to consume food today",
		Thesaurus{
			"big": {
				"large": struct{}{},
			},
			"eat": {
				"consume": struct{}{},
			},
		},
		false,
	},
}

func TestAreSentencesEquivalent(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := AreSentencesEquivalent(tc.sentence1, tc.sentence2, tc.thesaurus); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkAreSentencesEquivalent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			AreSentencesEquivalent(tc.sentence1, tc.sentence2, tc.thesaurus)
		}
	}
}

package day259

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	dict   []string
	result []rune
}{
	{[]string{"cat", "calf", "dog", "bear"}, []rune{'b'}},
}

func TestGhostWinningLetterPlayerOne(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := GhostWinningLetterPlayerOne(tc.dict); !reflect.DeepEqual(result, tc.result) {
			t.Errorf("Expected %v, got %v", string(tc.result), string(result))
		}
	}
}

func BenchmarkGhostWinningLetterPlayerOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			GhostWinningLetterPlayerOne(tc.dict)
		}
	}
}

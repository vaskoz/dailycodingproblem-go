package day250

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	first, second, result string
	solution              map[rune]int
}{
	{
		"SEND",
		"MORE",
		"MONEY",
		map[rune]int{
			'S': 7,
			'E': 5,
			'N': 3,
			'D': 1,
			'M': 0,
			'O': 8,
			'R': 2,
			'Y': 6,
		},
	},
	{
		"ABCDEFGHIJKL",
		"MNOPQRSTUVWXYZ",
		"TOOMANY",
		nil,
	},
}

func TestCryptarithmeticAdditionPuzzle(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := CryptarithmeticAdditionPuzzle(tc.first, tc.second, tc.result); !reflect.DeepEqual(result, tc.solution) {
			t.Errorf("Expected %v, got %v", tc.solution, result)
		}
	}
}

func BenchmarkCryptarithmeticAdditionPuzzle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CryptarithmeticAdditionPuzzle(tc.first, tc.second, tc.result)
		}
	}
}

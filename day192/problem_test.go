package day192

import "testing"

var testcases = []struct {
	advanceFrom []int
	possible    bool
}{
	{[]int{1, 3, 1, 2, 0, 1}, true},
	{[]int{1, 2, 1, 0, 0}, false},
	{[]int{0}, true},
	{[]int{}, true},
}

func TestCanAdvanceToEndBrute(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if possible := CanAdvanceToEndBrute(tc.advanceFrom); possible != tc.possible {
			t.Errorf("Given %v, expected %v, got %v", tc.advanceFrom, tc.possible, possible)
		}
	}
}

func BenchmarkCanAdvanceToEndBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CanAdvanceToEndBrute(tc.advanceFrom)
		}
	}
}

func TestCanAdvanceToEndLinear(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if possible := CanAdvanceToEndLinear(tc.advanceFrom); possible != tc.possible {
			t.Errorf("Given %v, expected %v, got %v", tc.advanceFrom, tc.possible, possible)
		}
	}
}

func BenchmarkCanAdvanceToEndLinear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CanAdvanceToEndLinear(tc.advanceFrom)
		}
	}
}

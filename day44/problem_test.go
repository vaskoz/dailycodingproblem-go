package day44

import "testing"

var testcases = []struct {
	input      []int
	inversions int
}{
	{[]int{2, 4, 1, 3, 5}, 3},
}

func TestInversionCountBrute(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := InversionCountBrute(tc.input); result != tc.inversions {
			t.Errorf("Expected %d but got %v", tc.inversions, result)
		}
	}
}

func BenchmarkInversionCountBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			InversionCountBrute(tc.input)
		}
	}
}

func TestInversionCount(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := InversionCount(tc.input); result != tc.inversions {
			t.Errorf("Expected %d but got %v", tc.inversions, result)
		}
	}
}

func BenchmarkInversionCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			InversionCount(tc.input)
		}
	}
}

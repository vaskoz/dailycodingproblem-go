package day44

import "testing"

var testcases = []struct {
	input      []int
	inversions int
}{
	{[]int{2, 4, 1, 3, 5}, 3},
	{[]int{100, 99, 98, 97, 96, 95, 94, 93, 92, 91, 90, 89, 88, 87, 86, 85, 84, 83, 82, 81, 80, 79, 78, 77, 76, 75, 74, 73, 72, 71, 70, 69, 68, 67, 66, 65, 64, 63, 62, 61, 60, 59, 58, 57, 56, 55, 54, 53, 52, 51, 50, 49, 48, 47, 46, 45, 44, 43, 42, 41, 40, 39, 38, 37, 36, 35, 34, 33, 32, 31, 30, 29, 28, 27, 26, 25}, 2850},
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

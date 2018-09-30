package day40

import "testing"

var testcases = []struct {
	input    []int
	expected int
}{
	{[]int{6, 1, 3, 3, 3, 6, 6}, 1},
	{[]int{13, 19, 13, 13}, 19},
}

func TestFindOnce(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := FindOnce(tc.input); tc.expected != result {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkFindOnce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			FindOnce(tc.input)
		}
	}
}

package day196

import "testing"

var testcases = []struct {
	head     *BinaryTree
	expected int
}{
	{&BinaryTree{
		5,
		&BinaryTree{2, nil, nil},
		&BinaryTree{-5, nil, nil},
	}, 2},
}

func TestMostFrequentSubtreeSum(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MostFrequentSubtreeSum(tc.head); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkMostFrequentSubtreeSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MostFrequentSubtreeSum(tc.head)
		}
	}
}

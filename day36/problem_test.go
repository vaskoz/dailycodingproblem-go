package day36

import "testing"

var testcases = []struct {
	root     *BST
	expected int
}{
	{&BST{10, nil, &BST{20, nil, &BST{30, nil, nil}}}, 20},
	{&BST{10, &BST{5, nil, &BST{8, nil, nil}}, nil}, 8},
}

func TestSecondLargest(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := SecondLargest(tc.root); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkSecondLargest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SecondLargest(tc.root)
		}
	}
}

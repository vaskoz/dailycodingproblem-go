package day89

import "testing"

var testcases = []struct {
	root    *BST
	isValid bool
}{
	{&BST{}, true},
	{&BST{10, &BST{1, nil, nil}, &BST{20, nil, nil}}, true},
	{&BST{10, &BST{20, nil, nil}, &BST{1, nil, nil}}, false},
}

func TestIsValidBST(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := IsValidBST(tc.root); result != tc.isValid {
			t.Errorf("Expected %v got %v", tc.isValid, result)
		}
	}
}

func BenchmarkIsValidBST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			IsValidBST(tc.root)
		}
	}
}

package day107

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	tree     *BinaryTree
	expected []int
}{
	{&BinaryTree{1,
		&BinaryTree{2, nil, nil},
		&BinaryTree{3, &BinaryTree{4, nil, nil}, &BinaryTree{5, nil, nil}}},
		[]int{1, 2, 3, 4, 5}},
	{&BinaryTree{1,
		&BinaryTree{3, &BinaryTree{4, nil, nil}, &BinaryTree{5, nil, nil}},
		&BinaryTree{2, nil, nil}},
		[]int{1, 3, 2, 4, 5}},
}

func TestBinaryTreeByLevel(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := BinaryTreeByLevel(tc.tree); !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkBinaryTreeByLevel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			BinaryTreeByLevel(tc.tree)
		}
	}
}

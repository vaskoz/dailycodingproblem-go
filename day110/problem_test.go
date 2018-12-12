package day110

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	head     *IntBinaryTree
	expected [][]int
}{
	{&IntBinaryTree{Value: 1,
		Left: &IntBinaryTree{2, nil, nil},
		Right: &IntBinaryTree{3,
			&IntBinaryTree{4, nil, nil},
			&IntBinaryTree{5, nil, nil}}},
		[][]int{{1, 2}, {1, 3, 4}, {1, 3, 5}}},
}

func TestAllPathsToLeaves(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := AllPathsToLeaves(tc.head); !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkAllPathsToLeaves(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			AllPathsToLeaves(tc.head)
		}
	}
}

package day471

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	n         int
	preorders [][]int
}{
	{
		3,
		[][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{3, 1, 2},
			{3, 2, 1},
		},
	},
}

func TestConstructAllPossibleBSTs(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		result := ConstructAllPossibleBSTs(tc.n)
		preorders := make([][]int, 0, len(result))

		for _, root := range result {
			preorders = append(preorders, preorder(root, []int{}))
		}

		if !reflect.DeepEqual(preorders, tc.preorders) {
			t.Errorf("Expected %v, got %v", tc.preorders, preorders)
		}
	}
}

func BenchmarkConstructAllPossibleBSTs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ConstructAllPossibleBSTs(tc.n)
		}
	}
}

func preorder(root *BST, lst []int) []int {
	if root != nil {
		lst = append(lst, root.Value)
		lst = preorder(root.Left, lst)
		lst = preorder(root.Right, lst)
	}

	return lst
}

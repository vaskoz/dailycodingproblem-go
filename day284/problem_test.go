package day284

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	root    *BinaryTree
	node    int
	cousins []int
}{
	{
		&BinaryTree{
			1,
			&BinaryTree{
				2,
				&BinaryTree{4, nil, nil},
				&BinaryTree{5, nil, nil},
			},
			&BinaryTree{
				3,
				nil,
				&BinaryTree{6, nil, nil},
			},
		},
		4,
		[]int{6},
	},
	{
		&BinaryTree{
			1,
			&BinaryTree{
				2,
				&BinaryTree{4, nil, nil},
				&BinaryTree{5, nil, nil},
			},
			&BinaryTree{
				3,
				nil,
				&BinaryTree{6, nil, nil},
			},
		},
		10,
		[]int{},
	},
	{
		&BinaryTree{
			1,
			&BinaryTree{
				2,
				&BinaryTree{4, nil, nil},
				&BinaryTree{5, nil, nil},
			},
			&BinaryTree{
				3,
				nil,
				&BinaryTree{6, nil, nil},
			},
		},
		6,
		[]int{4, 5},
	},
}

func TestFindCousins(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if cousins := FindCousins(tc.root, tc.node); !reflect.DeepEqual(cousins, tc.cousins) {
			t.Errorf("Expected %v, got %v", tc.cousins, cousins)
		}
	}
}

func BenchmarkFindCousins(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			FindCousins(tc.root, tc.node)
		}
	}
}

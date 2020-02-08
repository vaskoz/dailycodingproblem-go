package day442

import (
	"testing"
)

// nolint
var testcases = []struct {
	seq      []int
	expected *IntBinaryTree
}{
	{
		[]int{3, 2, 6, 1, 9},
		&IntBinaryTree{
			1,
			&IntBinaryTree{
				2,
				&IntBinaryTree{3, nil, nil},
				&IntBinaryTree{6, nil, nil},
			},
			&IntBinaryTree{9, nil, nil},
		},
	},
}

func TestCartesianTree(t *testing.T) {
	t.Parallel()

	for tcid, tc := range testcases {
		if result := CartesianTree(tc.seq); !equalTree(result, tc.expected) {
			t.Errorf("Trees do not match for TCID%d", tcid)
		}
	}
}

func BenchmarkCartesianTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CartesianTree(tc.seq)
		}
	}
}

func equalTree(a, b *IntBinaryTree) bool {
	if a == nil && b == nil {
		return true
	} else if a != nil && b != nil {
		return a.Value == b.Value &&
			equalTree(a.Left, b.Left) &&
			equalTree(a.Right, b.Right)
	}

	return false
}

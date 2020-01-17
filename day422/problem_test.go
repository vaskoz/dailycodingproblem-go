package day422

import "testing"

// nolint
var testcases = []struct {
	one, two, merged *BinaryTree
}{
	{
		&BinaryTree{
			5,
			&BinaryTree{
				Val: 9,
			},
			&BinaryTree{
				Val: 11,
			},
		},
		&BinaryTree{
			6,
			&BinaryTree{
				Val: 10,
			},
			&BinaryTree{
				Val: 21,
			},
		},
		&BinaryTree{
			11,
			&BinaryTree{
				Val: 19,
			},
			&BinaryTree{
				Val: 32,
			},
		},
	},
	{
		&BinaryTree{
			5,
			&BinaryTree{
				Val: 9,
			},
			nil,
		},
		&BinaryTree{
			6,
			nil,
			&BinaryTree{
				Val: 21,
			},
		},
		&BinaryTree{
			11,
			&BinaryTree{
				Val: 9,
			},
			&BinaryTree{
				Val: 21,
			},
		},
	},
	{
		&BinaryTree{
			5,
			&BinaryTree{
				Val: 9,
			},
			nil,
		},
		nil,
		&BinaryTree{
			5,
			&BinaryTree{
				Val: 9,
			},
			nil,
		},
	},
}

func TestMergeBinaryTree(t *testing.T) {
	t.Parallel()

	for tcid, tc := range testcases {
		if result := MergeBinaryTrees(tc.one, tc.two); !equal(result, tc.merged) {
			t.Errorf("Trees in TCID%d don't match", tcid)
		}
	}
}

func BenchmarkMergeBinaryTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MergeBinaryTrees(tc.one, tc.two)
		}
	}
}

func equal(a, b *BinaryTree) bool {
	switch {
	case a == nil && b != nil:
		return false
	case a != nil && b == nil:
		return false
	case a == nil && b == nil:
		return true
	case a.Val != b.Val:
		return false
	}

	return equal(a.Left, b.Left) && equal(a.Right, b.Right)
}

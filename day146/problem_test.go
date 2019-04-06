package day146

import "testing"

var testcases = []struct {
	head, expected *BinaryTree
}{
	{&BinaryTree{
		0,
		&BinaryTree{
			1,
			nil,
			nil,
		},
		&BinaryTree{
			0,
			&BinaryTree{
				1,
				&BinaryTree{
					0,
					nil,
					nil,
				},
				&BinaryTree{
					0,
					nil,
					nil,
				},
			},
			&BinaryTree{
				0,
				nil,
				nil,
			},
		},
	},
		&BinaryTree{
			0,
			&BinaryTree{1, nil, nil},
			&BinaryTree{0,
				&BinaryTree{1, nil, nil},
				nil},
		},
	},
}

func TestPruneZeroSubtrees(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		if result := PruneZeroSubtrees(tc.head); !treeEqual(result, tc.expected) {
			t.Errorf("TCID%d didn't match", tcid)
		}
	}
}

func BenchmarkPruneZeroSubtrees(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			PruneZeroSubtrees(tc.head)
		}
	}
}

func treeEqual(one, two *BinaryTree) bool {
	switch {
	case one == nil && two == nil:
		return true
	case one == nil && two != nil:
		return false
	case one != nil && two == nil:
		return false
	}
	return treeEqual(one.Left, two.Left) && treeEqual(one.Right, two.Right)
}

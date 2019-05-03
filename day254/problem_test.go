package day254

import "testing"

// nolint
var testcases = []struct {
	input, expected *BinaryTree
}{
	{
		&BinaryTree{
			0,
			&BinaryTree{
				1,
				&BinaryTree{
					3,
					nil,
					&BinaryTree{5, nil, nil},
				},
				nil,
			},
			&BinaryTree{
				2,
				nil,
				&BinaryTree{
					4,
					&BinaryTree{6, nil, nil},
					&BinaryTree{7, nil, nil},
				},
			},
		},
		&BinaryTree{
			0,
			&BinaryTree{5, nil, nil},
			&BinaryTree{
				4,
				&BinaryTree{6, nil, nil},
				&BinaryTree{7, nil, nil},
			},
		},
	},
}

func TestConvertBinaryTreeToFull(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		tree := copyTree(tc.input)
		if result := ConvertBinaryTreeToFull(tree); !equalTree(result, tc.expected) {
			t.Errorf("Trees don't match for tcid%d", tcid)
		}
	}
}

func copyTree(a *BinaryTree) *BinaryTree {
	if a == nil {
		return nil
	}
	n := &BinaryTree{}
	n.Value = a.Value
	n.Left = copyTree(a.Left)
	n.Right = copyTree(a.Right)
	return n
}

func equalTree(a, b *BinaryTree) bool {
	if a == nil && b == nil {
		return true
	}
	if a != nil && b != nil {
		return a.Value == b.Value &&
			equalTree(a.Left, b.Left) &&
			equalTree(a.Right, b.Right)
	}
	return false
}

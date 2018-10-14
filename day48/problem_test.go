package day48

import (
	"testing"
)

var testcases = []struct {
	preorder, inorder []string
	expected          *BinaryTree
}{
	{[]string{"a", "b", "d", "e", "c", "f", "g"},
		[]string{"d", "b", "e", "a", "f", "c", "g"},
		&BinaryTree{"a",
			&BinaryTree{"b", &BinaryTree{"d", nil, nil}, &BinaryTree{"e", nil, nil}},
			&BinaryTree{"c", &BinaryTree{"f", nil, nil}, &BinaryTree{"g", nil, nil}}}},
}

func TestReconstructBinaryTree(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result, _ := Reconstruct(tc.preorder, tc.inorder); !treesEqual(result, tc.expected) {
			t.Errorf("Expected both trees to be equal")
		}
	}
}

func BenchmarkReconstructBinaryTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			Reconstruct(tc.preorder, tc.inorder)
		}
	}
}

func treesEqual(a, b *BinaryTree) bool {
	if a == nil && b == nil {
		return true
	} else if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}
	if a.value != b.value {
		return false
	}
	if leftEqual := treesEqual(a.left, b.left); !leftEqual {
		return false
	} else if rightEqual := treesEqual(a.right, b.right); !rightEqual {
		return false
	}
	return true
}

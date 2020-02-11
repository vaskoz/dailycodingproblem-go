package day348

import (
	"testing"
)

// nolint
var testcases = []struct {
	words    []string
	tree     *TernarySearchTree
	notWords []string
}{
	{
		[]string{"code", "cob", "be", "ax", "war", "we"},
		&TernarySearchTree{
			'c',
			&TernarySearchTree{
				'b',
				&TernarySearchTree{
					'a',
					nil,
					&TernarySearchTree{
						'x', nil, nil, nil,
					},
					nil,
				},
				&TernarySearchTree{
					'e', nil, nil, nil,
				},
				nil,
			},
			&TernarySearchTree{
				'o',
				nil,
				&TernarySearchTree{
					'd',
					&TernarySearchTree{
						'b',
						nil,
						nil,
						nil,
					},
					&TernarySearchTree{
						'e', nil, nil, nil,
					},
					nil,
				},
				nil,
			},
			&TernarySearchTree{
				'w',
				nil,
				&TernarySearchTree{
					'a',
					nil,
					&TernarySearchTree{
						'r', nil, nil, nil,
					},
					&TernarySearchTree{
						'e', nil, nil, nil,
					},
				},
				nil,
			},
		},
		[]string{"foo", "bar", "baz", "blahblahblah"},
	},
}

func TestTernarySearchTree(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		tree := &TernarySearchTree{}

		for _, word := range tc.words {
			tree.Insert(word)
		}

		if !equalTernaryTree(tree, tc.tree) {
			t.Errorf("Expected trees to be equal but they aren't")
		}

		for _, word := range tc.words {
			if !tree.Search(word) {
				t.Errorf("Searching for word '%v' should return true", word)
			}
		}

		for _, word := range tc.notWords {
			if tree.Search(word) {
				t.Errorf("Searching for notWord '%v' should return false", word)
			}
		}
	}
}

func BenchmarkTernarySearchTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			tree := &TernarySearchTree{}

			for _, word := range tc.words {
				tree.Insert(word)
			}

			for _, word := range tc.words {
				tree.Search(word)
			}
		}
	}
}

func equalTernaryTree(a, b *TernarySearchTree) bool {
	if a == nil && b == nil {
		return true
	} else if a != nil && b != nil {
		return a.letter == b.letter &&
			equalTernaryTree(a.left, b.left) &&
			equalTernaryTree(a.middle, b.middle) &&
			equalTernaryTree(a.right, b.right)
	}

	return false
}

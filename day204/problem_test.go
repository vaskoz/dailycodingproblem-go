package day204

import "testing"

var testcases = []struct {
	tree  *CompleteBinaryTree
	count int
}{
	{&CompleteBinaryTree{
		"a",
		&CompleteBinaryTree{
			"b",
			&CompleteBinaryTree{
				"d",
				&CompleteBinaryTree{"h", nil, nil},
				&CompleteBinaryTree{"i", nil, nil},
			},
			&CompleteBinaryTree{
				"e",
				&CompleteBinaryTree{"j", nil, nil},
				&CompleteBinaryTree{"k", nil, nil},
			}},
		&CompleteBinaryTree{
			"c",
			&CompleteBinaryTree{
				"f",
				&CompleteBinaryTree{"l", nil, nil},
				nil,
			},
			&CompleteBinaryTree{"g", nil, nil},
		},
	}, 12},
}

func TestCountCompleteBinaryTree(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if count := CountCompleteBinaryTree(tc.tree); count != tc.count {
			t.Errorf("Expected %v, got %v", tc.count, count)
		}
	}
}

func BenchmarkCountCompleteBinaryTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CountCompleteBinaryTree(tc.tree)
		}
	}
}

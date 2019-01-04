package day135

import "testing"

var testcases = []struct {
	tree   *BinaryTree
	minSum int
}{
	{&BinaryTree{10,
		&BinaryTree{5, nil, &BinaryTree{2, nil, nil}},
		&BinaryTree{5, nil, &BinaryTree{1, &BinaryTree{-1, nil, nil}, nil}},
	}, 15},
	{&BinaryTree{10,
		&BinaryTree{5, nil, &BinaryTree{1, &BinaryTree{-1, nil, nil}, nil}},
		&BinaryTree{5, nil, &BinaryTree{2, nil, nil}},
	}, 15},
}

func TestMinPathSum(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MinPathSum(tc.tree); result != tc.minSum {
			t.Errorf("Expected %v got %v", tc.minSum, result)
		}
	}
}

func BenchmarkMinPathSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MinPathSum(tc.tree)
		}
	}
}

package day94

import "testing"

var tree = &BinaryTree{
	-15,
	&BinaryTree{5,
		&BinaryTree{-8,
			&BinaryTree{2, nil, nil},
			&BinaryTree{6, nil, nil}},
		&BinaryTree{1, nil, nil}},
	&BinaryTree{6,
		&BinaryTree{3, nil, nil},
		&BinaryTree{9, nil,
			&BinaryTree{0,
				&BinaryTree{4, nil, nil},
				&BinaryTree{-1, &BinaryTree{10, nil, nil}, nil},
			},
		},
	},
}

func TestMaxPathSum(t *testing.T) {
	t.Parallel()
	result := MaxPathSum(tree)
	if result != 27 {
		t.Errorf("Expected 27 got %v", result)
	}
}

func BenchmarkMaxPathSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxPathSum(tree)
	}
}

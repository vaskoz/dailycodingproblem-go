package day247

import "testing"

// nolint
var testcases = []struct {
	root           *BinaryTree
	heightBalanced bool
}{
	{
		&BinaryTree{
			&BinaryTree{nil, nil},
			&BinaryTree{nil, nil},
		},
		true,
	},
	{
		&BinaryTree{
			nil,
			&BinaryTree{nil, nil},
		},
		true,
	},
	{
		&BinaryTree{
			nil,
			&BinaryTree{&BinaryTree{nil, nil}, nil},
		},
		false,
	},
	{
		&BinaryTree{
			nil,
			&BinaryTree{
				&BinaryTree{
					nil,
					&BinaryTree{nil, nil},
				},
				nil,
			},
		},
		false,
	},
	{
		&BinaryTree{
			&BinaryTree{
				nil,
				&BinaryTree{
					&BinaryTree{nil, nil},
					nil,
				},
			},
			nil,
		},
		false,
	},
}

func TestIsHeightBalanced(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := IsHeightBalanced(tc.root); result != tc.heightBalanced {
			t.Errorf("expected %v, got %v", tc.heightBalanced, result)
		}
	}
}

func BenchmarkIsHeightBalanced(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			IsHeightBalanced(tc.root)
		}
	}
}

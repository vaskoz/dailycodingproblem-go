package day394

import "testing"

// nolint
var testcases = []struct {
	root     *IntBinaryTree
	target   int
	expected bool
}{
	{
		&IntBinaryTree{
			8,
			&IntBinaryTree{
				4,
				&IntBinaryTree{2, nil, nil},
				&IntBinaryTree{6, nil, nil},
			},
			&IntBinaryTree{
				13,
				nil,
				&IntBinaryTree{19, nil, nil},
			},
		},
		18,
		true,
	},
	{
		&IntBinaryTree{
			8,
			&IntBinaryTree{
				4,
				&IntBinaryTree{2, nil, nil},
				&IntBinaryTree{6, nil, nil},
			},
			&IntBinaryTree{
				13,
				nil,
				&IntBinaryTree{19, nil, nil},
			},
		},
		100,
		false,
	},
	{
		&IntBinaryTree{
			8,
			&IntBinaryTree{
				4,
				&IntBinaryTree{2, nil, nil},
				&IntBinaryTree{6, nil, nil},
			},
			&IntBinaryTree{
				13,
				nil,
				&IntBinaryTree{19, nil, nil},
			},
		},
		21,
		false,
	},
	{
		&IntBinaryTree{
			8,
			&IntBinaryTree{
				4,
				&IntBinaryTree{2, nil, nil},
				&IntBinaryTree{6, nil, nil},
			},
			&IntBinaryTree{
				13,
				nil,
				&IntBinaryTree{19, nil, nil},
			},
		},
		40,
		true,
	},
}

func TestRootToLeafTarget(t *testing.T) {
	t.Parallel()

	for tcid, tc := range testcases {
		if res := RootToLeafTarget(tc.root, tc.target); res != tc.expected {
			t.Errorf("TCID%d expected %v, got %v", tcid, tc.expected, res)
		}
	}
}

func BenchmarkRootToLeafTarget(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			RootToLeafTarget(tc.root, tc.target)
		}
	}
}

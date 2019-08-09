package day343

import "testing"

// nolint
var testcases = []struct {
	root     *BST
	a, b     int
	expected int
}{
	{
		&BST{
			5,
			&BST{
				3,
				&BST{2, nil, nil},
				&BST{4, nil, nil},
			},
			&BST{
				8,
				&BST{6, nil, nil},
				&BST{10, nil, nil},
			},
		},
		4, 9,
		23,
	},
	{
		&BST{
			5,
			&BST{
				3,
				&BST{2, nil, nil},
				&BST{4, nil, nil},
			},
			&BST{
				8,
				&BST{6, nil, nil},
				&BST{10, nil, nil},
			},
		},
		9, 4,
		0,
	},
}

func TestSumBSTRange(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if sum := SumBSTRange(tc.root, tc.a, tc.b); sum != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, sum)
		}
	}
}

func BenchmarkSumBSTRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SumBSTRange(tc.root, tc.a, tc.b)
		}
	}
}

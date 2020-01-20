package day426

import "testing"

// nolint
var testcases = []struct {
	tree             *BinaryTree
	minLevel, minSum int
}{
	{nil, 0, 0},
	{&BinaryTree{5,
		&BinaryTree{10, nil, nil},
		&BinaryTree{20, nil, nil}},
		1,
		5},
	{&BinaryTree{5,
		&BinaryTree{10, nil, nil},
		&BinaryTree{-20, nil, nil}},
		2,
		-10},
	{&BinaryTree{5,
		&BinaryTree{10,
			&BinaryTree{30,
				nil,
				&BinaryTree{0,
					&BinaryTree{-7, nil, nil},
					&BinaryTree{-9, nil, nil}}},
			nil},
		&BinaryTree{20,
			nil,
			&BinaryTree{0,
				&BinaryTree{3, nil, nil},
				&BinaryTree{5, nil, nil}}}},
		5,
		-16},
	{&BinaryTree{5,
		&BinaryTree{10,
			&BinaryTree{30,
				nil,
				&BinaryTree{-9, nil, nil}},
			nil},
		&BinaryTree{20,
			nil,
			&BinaryTree{0,
				&BinaryTree{3, nil, nil},
				&BinaryTree{5, nil, nil}}}},
		4,
		-1},
}

func TestMinimumSumLevel(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if sum, level := MinimumSumLevel(tc.tree); sum != tc.minSum && level != tc.minLevel {
			t.Errorf("Expected (%d,%d), but got (%d,%d)", tc.minSum, tc.minLevel, sum, level)
		}
	}
}

func BenchmarkMinimumSumLevel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MinimumSumLevel(tc.tree)
		}
	}
}

package day115

import "testing"

var testcases = []struct {
	t, s     *IntBinaryTree
	expected bool
}{
	{&IntBinaryTree{
		10,
		&IntBinaryTree{15, nil, nil},
		&IntBinaryTree{20, nil, nil}},
		&IntBinaryTree{15,
			&IntBinaryTree{10,
				&IntBinaryTree{15, nil, nil},
				&IntBinaryTree{20, nil, nil}},
			&IntBinaryTree{20, nil, nil}},
		true},
	{&IntBinaryTree{
		10, nil, nil},
		&IntBinaryTree{15,
			&IntBinaryTree{25,
				&IntBinaryTree{15, nil, nil},
				&IntBinaryTree{20, nil, nil}},
			&IntBinaryTree{20, nil, nil}},
		false},
	{&IntBinaryTree{10, nil, nil},
		&IntBinaryTree{15, nil, nil},
		false},
}

func TestSameTreeSubstructure(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		if result := SameTreeSubstructure(tc.t, tc.s); result != tc.expected {
			t.Errorf("TCID%d expected %v got %v", tcid, tc.expected, result)
		}
	}
}

func BenchmarkSameTreeSubstructure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SameTreeSubstructure(tc.t, tc.s)
		}
	}
}

package day8

import "testing"

func TestCountUnivalSubtrees(t *testing.T) {
	t.Parallel()
	root := &BinaryTree{value: 0,
		left: &BinaryTree{value: 1},
		right: &BinaryTree{value: 0,
			left: &BinaryTree{value: 1,
				left:  &BinaryTree{value: 1},
				right: &BinaryTree{value: 1}},
			right: &BinaryTree{value: 0}}}
	if count := CountUnivalSubtrees(root); count != 5 {
		t.Errorf("expected %v, but got %v", 5, count)
	}
	root = &BinaryTree{value: 0,
		left: &BinaryTree{value: 1},
		right: &BinaryTree{value: 0,
			left: &BinaryTree{value: 0},
			right: &BinaryTree{value: 1,
				left:  &BinaryTree{value: 1},
				right: &BinaryTree{value: 1}}}}
	if count := CountUnivalSubtrees(root); count != 5 {
		t.Errorf("expected %v, but got %v", 5, count)
	}
}

func BenchmarkCountUnivalSubtrees(b *testing.B) {
	root1 := &BinaryTree{value: 0,
		left: &BinaryTree{value: 1},
		right: &BinaryTree{value: 0,
			left: &BinaryTree{value: 1,
				left:  &BinaryTree{value: 1},
				right: &BinaryTree{value: 1}},
			right: &BinaryTree{value: 0}}}
	root2 := &BinaryTree{value: 0,
		left: &BinaryTree{value: 1},
		right: &BinaryTree{value: 0,
			left: &BinaryTree{value: 0},
			right: &BinaryTree{value: 1,
				left:  &BinaryTree{value: 1},
				right: &BinaryTree{value: 1}}}}
	for i := 0; i < b.N; i++ {
		CountUnivalSubtrees(root1)
		CountUnivalSubtrees(root2)
	}
}

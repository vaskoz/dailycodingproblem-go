package day116

import "testing"

func TestGenerateUnboundBinaryTree(t *testing.T) {
	t.Parallel()
	tree := GenerateUnboundBinaryTree()
	for i := 0; i < 1000; i++ {
		tree = tree.Left()
		tree = tree.Right()
	}
	if tree.Value() != UnboundBinaryTree(nil) {
		t.Errorf("The values are always nil of the interface type")
	}
	if tree.Left() == nil {
		t.Errorf("The tree should be infinite. Left is nil")
	}
	if tree.Right() == nil {
		t.Errorf("The tree should be infinite. Right is nil")
	}
}

func BenchmarkGenerateUnboundBinaryTree(b *testing.B) {
	tree := GenerateUnboundBinaryTree()
	for i := 0; i < b.N; i++ {
		tree = tree.Left()
		tree = tree.Right()
	}
}

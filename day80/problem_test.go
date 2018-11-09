package day80

import "testing"

// nodes = 5 depth = 3
func buildTestTree() *BinaryTree {
	return &BinaryTree{&BinaryTree{&BinaryTree{nil, nil},
		&BinaryTree{&BinaryTree{nil, nil}, nil}},
		&BinaryTree{nil, &BinaryTree{nil, nil}}}
}

func TestDeepest(t *testing.T) {
	t.Parallel()
	root := buildTestTree()
	if _, count := Deepest(root); count != 4 {
		t.Errorf("Expected 4 but got %v", count)
	}
}

func BenchmarkDeepest(b *testing.B) {
	root := buildTestTree()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Deepest(root)
	}
}

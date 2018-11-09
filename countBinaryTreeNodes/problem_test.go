package deepest

import "testing"

// nodes = 5 depth = 3
func buildTestTree() *BinaryTree {
	return &BinaryTree{&BinaryTree{&BinaryTree{nil, nil},
		&BinaryTree{&BinaryTree{nil, nil}, nil}},
		&BinaryTree{nil, &BinaryTree{nil, nil}}}
}

func TestCountNodes(t *testing.T) {
	t.Parallel()
	root := buildTestTree()
	if count := CountNodes(root); count != 7 {
		t.Errorf("Expected 7 but got %v", count)
	}
}

func BenchmarkCountNodes(b *testing.B) {
	root := buildTestTree()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountNodes(root)
	}
}

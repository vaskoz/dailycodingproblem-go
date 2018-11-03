package day72

import "testing"

func TestLargestPathValue(t *testing.T) {
	t.Parallel()
	nodes := []rune("ABACA")
	edges := []Edge{{0, 1}, {0, 2}, {2, 3}, {3, 4}}
	expected := 3
	if result, _ := LargestPathValue(nodes, edges); result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestLargestPathValueInfiniteLoop(t *testing.T) {
	t.Parallel()
	nodes := []rune("A")
	edges := []Edge{{0, 0}}
	if _, err := LargestPathValue(nodes, edges); err != ErrInfiniteLoop() {
		t.Errorf("Expected an error")
	}
}

func BenchmarkLargestPathValue(b *testing.B) {
	nodes := []rune("ABACA")
	edges := []Edge{{0, 1}, {0, 2}, {2, 3}, {3, 4}}
	for i := 0; i < b.N; i++ {
		LargestPathValue(nodes, edges)
	}
}

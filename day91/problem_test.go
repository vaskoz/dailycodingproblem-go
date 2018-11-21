package day91

import (
	"reflect"
	"testing"
)

func TestBadSnippet(t *testing.T) {
	t.Parallel()
	expected := []int{9, 9, 9, 9, 9, 9, 9, 9, 9}
	result := BadSnippet()
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v got %v", expected, result)
	}
}

func BenchmarkBadSnippet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BadSnippet()
	}
}

func TestFixedSnippet(t *testing.T) {
	t.Parallel()
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	result := FixedSnippet()
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v got %v", expected, result)
	}
}

func BenchmarkFixedSnippet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FixedSnippet()
	}
}

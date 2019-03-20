package day210

import (
	"testing"
)

func TestLongestCollatzSequence(t *testing.T) {
	t.Parallel()
	t.Skip("unskip test to verify result. takes too long to run")
	n, length := LongestCollatzSequence(1000000)
	if n != 837799 || length != 524 {
		t.Errorf("Expected (837799,524), got (%v,%v)", n, length)
	}
}

func TestLongestCollatzSequenceSmaller(t *testing.T) {
	t.Parallel()
	n, length := LongestCollatzSequence(10)
	if n != 9 || length != 19 {
		t.Errorf("Expected (9, 19), got (%v,%v)", n, length)
	}
}

func BenchmarkLongestCollatzSequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LongestCollatzSequence(10)
	}
}

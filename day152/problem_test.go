package day152

import (
	"testing"
)

func TestRandom(t *testing.T) {
	t.Parallel()
	r := NewRandom([]interface{}{"a", "b", "c"}, []float64{0.1, 0.8, 0.1})
	counts := make(map[interface{}]int)
	for i := 0; i < 10000; i++ {
		counts[r.Get()]++
	}
	if ac := counts["a"]; ac < 200 || ac > 1800 {
		t.Errorf("Expected approximately 10 percent of 'a'")
	}
	if bc := counts["b"]; bc < 7000 || bc > 9000 {
		t.Errorf("Expected approximately 80 percent of 'b'")
	}
	if cc := counts["c"]; cc < 200 || cc > 1800 {
		t.Errorf("Expected approximately 10 percent of 'c'")
	}
}

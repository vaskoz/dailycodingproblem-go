package day120

import "testing"

func TestGetInstance(t *testing.T) {
	t.Parallel()
	g := New()
	first := g.GetInstance()
	second := g.GetInstance()
	if first == second {
		t.Errorf("The first and second instance should not be the same")
	}
	for i := 0; i < 100; i++ {
		if first != g.GetInstance() {
			t.Errorf("First doesn't match")
		}
		if second != g.GetInstance() {
			t.Errorf("Second doesn't match")
		}
	}
}

func BenchmarkGetInstance(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.GetInstance()
	}
}

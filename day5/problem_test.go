package day5

import "testing"

func TestCar(t *testing.T) {
	t.Parallel()
	result := Car(Cons(3, 4))
	if result != 3 {
		t.Errorf("Expected 3 but got %d", result)
	}
}

func TestCdr(t *testing.T) {
	t.Parallel()
	result := Cdr(Cons(3, 4))
	if result != 4 {
		t.Errorf("Expected 4 but got %d", result)
	}
}

// Pretty much just benchmarking function call performance.
// Benchmarks just included for completeness, but not necessary.

func BenchmarkCar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Car(Cons(3, 4))
	}
}

func BenchmarkCdr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cdr(Cons(3, 4))
	}
}

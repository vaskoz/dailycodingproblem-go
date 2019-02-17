package day178

import "testing"

func TestFirstGame(t *testing.T) {
	t.Parallel()
	var d Dice
	if cost := FirstGame(&d); cost > 500 {
		t.Errorf("Expected result to be around 36, got %v", cost)
	}
}

func TestSecondGame(t *testing.T) {
	t.Parallel()
	var d Dice
	if cost := SecondGame(&d); cost > 500 {
		t.Errorf("Expected result to be around 36, got %v", cost)
	}
}

func BenchmarkFirstGame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var d Dice
		FirstGame(&d)
	}
}

func BenchmarkSecondGame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var d Dice
		SecondGame(&d)
	}
}

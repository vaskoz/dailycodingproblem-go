package day66

import "testing"

func TestFairTosser(t *testing.T) {
	t.Parallel()
	bc := NewBiasedCoin(99)
	biasedResults := make([]int, 2)
	for i := 0; i < 100000; i++ {
		biasedResults[bc.Toss()]++
	}
	if biasedResults[1] == 0 {
		t.Errorf("Should have received 1 at least once")
	}
	if ratio := biasedResults[0] / biasedResults[1]; ratio < 90 {
		t.Errorf("Expected the ratio of 0's to 1's to be near 99, but got %v", ratio)
	}
	ft := NewFairTosser(bc)
	fairResults := make([]int, 2)
	for i := 0; i < 100000; i++ {
		fairResults[ft.Toss()]++
	}
	if ratio := float64(fairResults[0]) / float64(fairResults[1]); ratio < 0.9 || ratio > 1.1 {
		t.Errorf("Expected the ratio of 0's to 1's to be near 1.0, but got %v", ratio)
	}
}

func BenchmarkFairTosser(b *testing.B) {
	bc := NewBiasedCoin(99)
	ft := NewFairTosser(bc)
	for i := 0; i < b.N; i++ {
		ft.Toss()
	}
}

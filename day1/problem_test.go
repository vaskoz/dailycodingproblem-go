package day1

import "testing"

func TestTwoSums(t *testing.T) {
	ten := []int{100, 75, 50, 25, 11, 45, 78, 99, 101, 12}
	tenK := 201
	t.Run("Brute", func(t *testing.T) {
		t.Parallel()
		if result := TwoSumBrute(ten, tenK); !result {
			t.Error("tenK should be true")
		}
		if result := TwoSumBrute(ten, 202); result {
			t.Error("tenK should be false")
		}
	})
	t.Run("Better", func(t *testing.T) {
		t.Parallel()
		if result := TwoSumBetter(ten, tenK); !result {
			t.Error("tenK should be true")
		}
		if result := TwoSumBetter(ten, 202); result {
			t.Error("tenK should be false")
		}
	})
	t.Run("Best", func(t *testing.T) {
		t.Parallel()
		if result := TwoSumBest(ten, tenK); !result {
			t.Error("tenK should be true")
		}
		if result := TwoSumBest(ten, 202); result {
			t.Error("tenK should be false")
		}
	})
}

func BenchmarkTwoSums(b *testing.B) {
	ten := []int{100, 75, 50, 25, 11, 45, 78, 99, 101, 12}
	tenK := 201
	b.Run("Brute", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			TwoSumBrute(ten, tenK)
			TwoSumBrute(ten, 202)
		}
	})
	b.Run("Better", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			TwoSumBetter(ten, tenK)
			TwoSumBetter(ten, 202)
		}
	})
	b.Run("Best", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			TwoSumBest(ten, tenK)
			TwoSumBest(ten, 202)
		}
	})
}

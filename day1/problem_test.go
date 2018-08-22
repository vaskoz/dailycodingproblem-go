package day1

import "testing"

func TestTwoSumsSmall(t *testing.T) {
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

func TestTwoSumsLarge(t *testing.T) {
	large := make([]int, 1000)
	largeK := 2
	for i := 1; i < 501; i++ {
		for j := 0; j < 2; j++ {
			large[2*(i-1)+j] = i
		}
	}
	t.Run("Brute", func(t *testing.T) {
		t.Parallel()
		if result := TwoSumBrute(large, largeK); !result {
			t.Error("tenK should be true")
		}
		if result := TwoSumBrute(large, 1); result {
			t.Error("tenK should be false")
		}
	})
	t.Run("Better", func(t *testing.T) {
		t.Parallel()
		if result := TwoSumBetter(large, largeK); !result {
			t.Error("tenK should be true")
		}
		if result := TwoSumBetter(large, 1); result {
			t.Error("tenK should be false")
		}
	})
	t.Run("Best", func(t *testing.T) {
		t.Parallel()
		if result := TwoSumBest(large, largeK); !result {
			t.Error("tenK should be true")
		}
		if result := TwoSumBest(large, 1); result {
			t.Error("tenK should be false")
		}
	})
}

func BenchmarkTwoSumsSmall(b *testing.B) {
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

func BenchmarkTwoSumsLarge(b *testing.B) {
	large := make([]int, 1000)
	largeK := 2
	for i := 1; i < 501; i++ {
		for j := 0; j < 2; j++ {
			large[2*(i-1)+j] = i
		}
	}
	b.Run("Brute", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			TwoSumBrute(large, largeK)
			TwoSumBrute(large, 1)
		}
	})
	b.Run("Better", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			TwoSumBetter(large, largeK)
			TwoSumBetter(large, 1)
		}
	})
	b.Run("Best", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			TwoSumBest(large, largeK)
			TwoSumBest(large, 1)
		}
	})
}

package day71

import "testing"

func TestRand5(t *testing.T) {
	t.Parallel()
	counter := make([]int, 6) // we ignore 0, because 1-7 inclusive.
	for i := 0; i < 1000000; i++ {
		counter[Rand5()]++
	}
	uniform := 1000000 / 5
	tolerance := 1500
	for i := 1; i <= 5; i++ {
		if counter[i] < uniform-tolerance || counter[i] > uniform+tolerance {
			t.Errorf("Expected closer to uniform %d got %d", uniform, counter[i])
		}
	}
}
func BenchmarkRand5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rand5()
	}
}

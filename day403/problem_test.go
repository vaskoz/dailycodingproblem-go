package day45

import "testing"

func TestRand7(t *testing.T) {
	t.Parallel()

	counter := make([]int, 8) // we ignore 0, because 1-7 inclusive.
	for i := 0; i < 1000000; i++ {
		counter[Rand7()]++
	}

	uniform := 1000000 / 7
	tolerance := 1500

	for i := 1; i <= 7; i++ {
		if counter[i] < uniform-tolerance || counter[i] > uniform+tolerance {
			t.Errorf("Expected closer to uniform %d got %d", uniform, counter[i])
		}
	}
}

func BenchmarkRand7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rand7()
	}
}

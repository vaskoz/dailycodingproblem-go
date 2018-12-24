package day124

import "testing"

var testcases = []struct {
	coins, rounds int
}{
	{100, 7},
}

func TestFairCoinFlipDownToOne(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if rounds := FairCoinFlipDownToOne(tc.coins); rounds != tc.rounds {
			t.Errorf("For %v coins, expected %v rounds got %v", tc.coins, tc.rounds, rounds)
		}
	}
}

func BenchmarkFairCoinFlipDownToOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			FairCoinFlipDownToOne(tc.coins)
		}
	}
}

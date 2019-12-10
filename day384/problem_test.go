package day384

import "testing"

// nolint
var testcases = []struct {
	denom         []int
	target        int
	expectedCoins int
	expectedErr   error
}{
	{[]int{1, 5, 10}, 56, 7, nil},
	{[]int{5, 8}, 15, 3, nil},
	{[]int{5, 10}, 2, 0, errImpossibleTarget},
}

func TestFewestCoinsBrute(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if res, err := FewestCoinsBrute(tc.denom, tc.target); res != tc.expectedCoins || err != tc.expectedErr {
			t.Errorf("Expected (%v,%v), got (%v,%v)", tc.expectedCoins, tc.expectedErr, res, err)
		}
	}
}

func BenchmarkFewestCoinsBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			FewestCoinsBrute(tc.denom, tc.target) // nolint
		}
	}
}

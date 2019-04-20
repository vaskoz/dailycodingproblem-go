package day240

import "testing"

var testcases = []struct {
	pairs    []int
	minSwaps int
}{
	{[]int{0, 2, 1, 2, 0, 1}, 2},
	{[]int{0, 3, 1, 1, 2, 2, 3, 0}, 1},
	{[]int{0, 3, 2, 3, 0, 2}, 2},
	{[]int{1, 2, 3, 3, 1, 2}, 1},
	{[]int{3, 2, 3, 1, 1, 2}, 2},
	{[]int{0, 3, 0, 2, 3, 2}, 2},
	{[]int{0, 3, 2, 3, 0, 2}, 2},
}

func TestMinSwapsAdjacentPairs(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		copied := make([]int, len(tc.pairs))
		copy(copied, tc.pairs)
		if result := MinSwapsAdjacentPairs(copied); result != tc.minSwaps {
			t.Errorf("Expected %v, got %v", tc.minSwaps, result)
		}
	}
}

func TestMinSwapsAdjacentPairsNotEventInput(t *testing.T) {
	t.Parallel()
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected an error for odd number of pairs")
		}
	}()
	pairs := []int{1, 1, 1}
	MinSwapsAdjacentPairs(pairs)
}

func TestMinSwapsAdjacentPairsMoreThanTwoInPair(t *testing.T) {
	t.Parallel()
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected an error if there are more than 2 in a pair")
		}
	}()
	pairs := []int{1, 1, 1, 1}
	MinSwapsAdjacentPairs(pairs)
}

func TestMinSwapsAdjacentPairsThreeInPair(t *testing.T) {
	t.Parallel()
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected an error if there are more than 2 in a pair")
		}
	}()
	pairs := []int{3, 2, 3, 3, 1, 1}
	MinSwapsAdjacentPairs(pairs)
}

func BenchmarkMinSwapsAdjacentPairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			copied := make([]int, len(tc.pairs))
			copy(copied, tc.pairs)
			MinSwapsAdjacentPairs(copied)
		}
	}
}

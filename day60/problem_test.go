package day60

import "testing"

var testcases = []struct {
	multiset []int
	expected bool
}{
	{[]int{15, 5, 20, 10, 35, 15, 10}, true},
	{[]int{15, 5, 20, 10, 35}, false},
	{[]int{5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2,
		5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5,
		2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2,
		5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5, 2}, true},
	{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, true},
}

func TestBruteForcePartitionSet(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := BruteForcePartitionSet(tc.multiset); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkBruteForcePartitionSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			BruteForcePartitionSet(tc.multiset)
		}
	}
}

func TestPartitionSet(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := PartitionSet(tc.multiset); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkPartitionSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			PartitionSet(tc.multiset)
		}
	}
}

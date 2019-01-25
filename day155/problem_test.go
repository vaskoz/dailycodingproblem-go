package day155

import "testing"

var testcases = []struct {
	nums     []int
	majority int
	err      error
}{
	{[]int{1, 2, 1, 1, 3, 1, 4, 0}, 1, nil},
	{[]int{}, 0, ErrNoMajority()},
	{[]int{4, 1, 2, 4, 1, 4, 1, 3, 4, 0, 4}, 4, nil},
	{[]int{1, 2, 4, 1, 4, 1, 4, 3, 4, 0, 4}, 4, nil},
}

func TestMajorityElementMap(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result, err := MajorityElementMap(tc.nums); err != tc.err || result != tc.majority {
			t.Errorf("Expected (%v,%v) got (%v,%v)", tc.majority, tc.err, result, err)
		}
	}
}

func TestMajorityElementSort(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result, err := MajorityElementSort(tc.nums); err != tc.err || result != tc.majority {
			t.Errorf("Expected (%v,%v) got (%v,%v)", tc.majority, tc.err, result, err)
		}
	}
}

func TestMajorityBoyerMoore(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if tc.err != nil {
			continue
		}
		if result, err := MajorityBoyerMoore(tc.nums); err != tc.err || result != tc.majority {
			t.Errorf("Expected (%v,%v) got (%v,%v)", tc.majority, tc.err, result, err)
		}
	}
}

func BenchmarkMajorityElementMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MajorityElementMap(tc.nums)
		}
	}
}

func BenchmarkMajorityElementSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MajorityElementSort(tc.nums)
		}
	}
}

func BenchmarkMajorityBoyerMoore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MajorityBoyerMoore(tc.nums)
		}
	}
}

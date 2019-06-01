package day282

import "testing"

// nolint
var testcases = []struct {
	nums    []int
	a, b, c int
	err     error
}{
	{
		[]int{10, 11, 12},
		0, 0, 0,
		ErrNoAnswer(),
	},
	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		3, 4, 5,
		nil,
	},
	{
		[]int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		3, 4, 5,
		nil,
	},
}

func TestPythagoreanTripletBrute(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if a, b, c, err := PythagoreanTripletBrute(tc.nums); a != tc.a || b != tc.b || c != tc.c || err != tc.err {
			t.Errorf("Expected (%v,%v,%v,%v), got (%v,%v,%v,%v)", tc.a, tc.b, tc.c, tc.err, a, b, c, err)
		}
	}
}

func BenchmarkPythagoreanTripletBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			PythagoreanTripletBrute(tc.nums) // nolint
		}
	}
}

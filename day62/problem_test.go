package day62

import "testing"

var testcases = []struct {
	n, m, expected int
}{
	{2, 2, 2},
	{5, 5, 70},
}

func TestCountPathsBrute(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := CountPathsBrute(tc.m, tc.n); result != tc.expected {
			t.Errorf("Expected %d got %d", tc.expected, result)
		}
	}
}

func TestCountPathsBruteBadArgs(t *testing.T) {
	t.Parallel()
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic")
		}
	}()
	CountPathsBrute(0, 0)
}

func BenchmarkCountPathsBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CountPathsBrute(tc.m, tc.n)
		}
	}
}

func TestCountPathsDP(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := CountPathsDP(tc.m, tc.n); result != tc.expected {
			t.Errorf("Expected %d got %d", tc.expected, result)
		}
	}
}

func TestCountPathsDPBadArgs(t *testing.T) {
	t.Parallel()
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic")
		}
	}()
	CountPathsDP(0, 0)
}

func BenchmarkCountPathsDP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CountPathsDP(tc.m, tc.n)
		}
	}
}

func TestCountPathsCalculate(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := CountPathsCalculate(tc.m, tc.n); result != tc.expected {
			t.Errorf("Expected %d got %d", tc.expected, result)
		}
	}
}

func TestCountPathsCalculateBadArgs(t *testing.T) {
	t.Parallel()
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic")
		}
	}()
	CountPathsCalculate(0, 0)
}

func BenchmarkCountPathsCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CountPathsCalculate(tc.m, tc.n)
		}
	}
}

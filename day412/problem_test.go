package day412

import "testing"

// nolint
var testcases = []struct {
	n      int
	result string
}{
	{1, "1"},
	{2, "11"},
	{3, "21"},
	{4, "1211"},
	{5, "111221"},
	{6, "312211"},
	{7, "13112221"},
	{8, "1113213211"},
}

func TestLookAndSay(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := LookAndSay(tc.n); result != tc.result {
			t.Errorf("Expected %v, got %v", tc.result, result)
		}
	}
}

func TestLookAndSayBadInput(t *testing.T) {
	t.Parallel()

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic for input <1")
		}
	}()
	LookAndSay(0)
}

func BenchmarkLookAndSay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LookAndSay(tc.n)
		}
	}
}

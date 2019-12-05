package day381

import "testing"

// nolint
var testcases = []struct {
	hexs     string
	expected string
}{
	{"deadbeef", "3q2+7w=="},
}

func TestBase64Delegate(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := Base64Delegate(tc.hexs); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, tc.hexs)
		}
	}
}

func TestBase64DelegateBadInput(t *testing.T) {
	t.Parallel()

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("should have paniced with bad input")
		}
	}()
	Base64Delegate("thisisnothex")
}

func BenchmarkBase64Delegate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			Base64Delegate(tc.hexs)
		}
	}
}

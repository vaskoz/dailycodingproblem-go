package day382

import "testing"

// nolint
var testcases = []struct {
	base64  string
	decoded string
}{
	{"3q2+7w==", "deadbeef"},
}

func TestBase64DecodeDelegate(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if res := Base64DecodeDelegate(tc.base64); res != tc.decoded {
			t.Errorf("Expected %v, got %v", tc.decoded, res)
		}
	}
}

func TestBase64DecodeDelegateBadInput(t *testing.T) {
	t.Parallel()

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("expected a panic for non-base64 input")
		}
	}()

	Base64DecodeDelegate("base64NOT[]")
}

func BenchmarkBase64DecodeDelegate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			Base64DecodeDelegate(tc.base64)
		}
	}
}

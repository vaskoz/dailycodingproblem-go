package day385

import "testing"

// nolint
var testcases = []struct {
	encoded, decoded string
}{
	{
		"7a575e5e5d12455d405e561254405d5f1276535b5e4b12715d565b5c551262405d505e575f",
		"Hello world from Daily Coding Problem",
	},
}

func TestDecryptSecretXorMessage(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if res := DecryptSecretXorMessage(tc.encoded); res != tc.decoded {
			t.Errorf("Expected %v, got %v", tc.decoded, res)
		}
	}
}

func TestDecryptSecretXorMessageBadHex(t *testing.T) {
	t.Parallel()

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic from non-hex input")
		}
	}()

	DecryptSecretXorMessage("nothex")
}

func BenchmarkDecryptSecretXorMessage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			DecryptSecretXorMessage(tc.encoded)
		}
	}
}

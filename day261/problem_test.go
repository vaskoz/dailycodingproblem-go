package day261

import "testing"

// nolint
var testcases = []struct {
	freq    map[rune]int
	decoded []string
	encoded []string
}{
	{
		freq: map[rune]int{
			'c': 1,
			's': 2,
			'a': 10,
			't': 11,
		},
		decoded: []string{"cats"},
		encoded: []string{"110100111"},
	},
}

func TestVasko(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		NewHuffman(tc.freq)
	}
}

func TestHuffmanEncode(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		h := NewHuffman(tc.freq)
		for i := range tc.decoded {
			if encode, err := h.Encode(tc.decoded[i]); encode != tc.encoded[i] {
				t.Errorf("Given %v, expected %v, got (%v,%v)", tc.decoded[i], tc.encoded[i], encode, err)
			}
		}
	}
}

func TestHuffmanDecode(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		h := NewHuffman(tc.freq)
		for i := range tc.encoded {
			if decode, _ := h.Decode(tc.encoded[i]); decode != tc.decoded[i] {
				t.Errorf("Given %v, expected %v, got %v", tc.encoded[i], tc.decoded[i], decode)
			}
		}
	}
}

func BenchmarkHuffman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			h := NewHuffman(tc.freq)
			for i := range tc.decoded {
				h.Encode(tc.decoded[i]) // nolint
				h.Decode(tc.encoded[i]) // nolint
			}
		}
	}
}

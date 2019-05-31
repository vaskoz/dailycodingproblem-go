package day261

import "testing"

// nolint
var testcases = []struct {
	freq       map[rune]int
	decoded    []string
	decodedErr []error
	encoded    []string
	encodedErr []error
}{
	{
		freq: map[rune]int{
			'c': 1,
			's': 2,
			'a': 10,
			't': 11,
		},
		decoded:    []string{"cats"},
		decodedErr: []error{nil},
		encoded:    []string{"110100111"},
		encodedErr: []error{nil},
	},
	{
		freq: map[rune]int{
			'c': 1,
			's': 2,
			'a': 10,
			't': 11,
		},
		decoded:    []string{"dogs"},
		decodedErr: []error{ErrTranslate()},
		encoded:    []string{"dogs"},
		encodedErr: []error{ErrTranslate()},
	},
}

func TestHuffmanEncode(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		h := NewHuffman(tc.freq)
		for i := range tc.decoded {
			if encode, err := h.Encode(tc.decoded[i]); err != tc.encodedErr[i] || encode != tc.encoded[i] {
				t.Errorf("Given %v, expected (%v,%v), got (%v,%v)", tc.decoded[i], tc.encoded[i], tc.encodedErr[i], encode, err)
			}
		}
	}
}

func TestHuffmanDecode(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		h := NewHuffman(tc.freq)
		for i := range tc.encoded {
			if decode, err := h.Decode(tc.encoded[i]); err != tc.decodedErr[i] || decode != tc.decoded[i] {
				t.Errorf("Given %v, expected (%v,%v), got (%v,%v)", tc.encoded[i], tc.decoded[i], tc.decodedErr[i], decode, err)
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

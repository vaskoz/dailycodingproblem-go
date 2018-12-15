package day114

import "testing"

var testcases = []struct {
	input, expected string
}{
	{"hello/world:here", "here/world:hello"},
	{"hello/world:here/", "/here:world/hello"},
	{"hello//world:here", "here/world/:hello"},
	{"/hello//world:here/", "/here/world/:hello/"},
}

func TestReverseWordsMaintainDelimters(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := ReverseWordsMaintainDelimters(tc.input); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkReverseWordsMaintainDelimters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ReverseWordsMaintainDelimters(tc.input)
		}
	}
}

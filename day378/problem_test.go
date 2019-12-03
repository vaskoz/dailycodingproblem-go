package day378

import "testing"

// nolint
var testcases = []struct {
	input    interface{}
	expected string
}{
	{
		[]interface{}{nil, 123, []string{"a", "b"}, map[string]string{"c": "d"}},
		`[null,123,["a","b"],{"c":"d"}]`,
	},
	{
		func() {},
		"",
	},
}

func TestToJSON(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := ToJSON(tc.input); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkToJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ToJSON(tc.input)
		}
	}
}

package day182

import "testing"

var testcases = []struct {
	g        UndirectedGraph
	expected bool
}{
	{UndirectedGraph{
		"a": {},
		"b": {},
	}, false},
	{UndirectedGraph{
		"a": {"b": struct{}{}},
		"b": {"a": struct{}{}},
	}, true},
	{UndirectedGraph{
		"a": {"b": struct{}{}},
		"b": {
			"a": struct{}{},
			"c": struct{}{},
		},
		"c": {"b": struct{}{}},
	}, true},
	{UndirectedGraph{
		"a": {
			"b": struct{}{},
			"c": struct{}{},
		},
		"b": {
			"a": struct{}{},
			"c": struct{}{},
		},
		"c": {
			"a": struct{}{},
			"b": struct{}{},
		},
	}, false},
}

func TestIsMinimallyConnected(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		if result := IsMinimallyConnected(tc.g); result != tc.expected {
			t.Errorf("TCID%d Expected %v, got %v", tcid, tc.expected, result)
		}
	}
}

func BenchmarkIsMinimallyConnected(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			IsMinimallyConnected(tc.g)
		}
	}
}

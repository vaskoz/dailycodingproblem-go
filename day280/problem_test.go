package day280

import "testing"

// nolint
var testcases = []struct {
	g      UndirectedGraph
	cyclic bool
}{
	{
		UndirectedGraph{
			0: {2: struct{}{}, 3: struct{}{}},
			1: {0: struct{}{}},
			2: {0: struct{}{}},
			3: {4: struct{}{}},
		},
		true,
	},
	{
		UndirectedGraph{
			0: {1: struct{}{}},
			1: {2: struct{}{}},
			2: {3: struct{}{}},
			3: {4: struct{}{}},
		},
		false,
	},
}

func TestHasCycle(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if cyclic := HasCycle(tc.g); cyclic != tc.cyclic {
			t.Errorf("Expected %v, got %v", tc.cyclic, cyclic)
		}
	}
}

func BenchmarkHasCycle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			HasCycle(tc.g)
		}
	}
}

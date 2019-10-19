package day344

import "testing"

// nolint
var testcases = []struct {
	tree     Tree
	expected int
}{
	{
		Tree{
			1: map[int]struct{}{
				2: struct{}{},
				3: struct{}{},
			},
			3: map[int]struct{}{
				4: struct{}{},
				5: struct{}{},
			},
			4: map[int]struct{}{
				6: struct{}{},
				7: struct{}{},
				8: struct{}{},
			},
		},
		2,
	},
	{
		Tree{
			0: map[int]struct{}{
				2: struct{}{},
				4: struct{}{},
				1: struct{}{},
			},
			2: map[int]struct{}{
				3: struct{}{},
			},
			4: map[int]struct{}{
				5: struct{}{},
			},
			5: map[int]struct{}{
				6: struct{}{},
				7: struct{}{},
			},
		},
		2,
	},
}

func TestMaxRemovedEdgesEvenSubtrees(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := MaxRemovedEdgesEvenSubtrees(tc.tree); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkMaxRemovedEdgesEvenSubtrees(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MaxRemovedEdgesEvenSubtrees(tc.tree)
		}
	}
}

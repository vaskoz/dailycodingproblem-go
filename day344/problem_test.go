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
				2: {},
				3: {},
			},
			3: map[int]struct{}{
				4: {},
				5: {},
			},
			4: map[int]struct{}{
				6: {},
				7: {},
				8: {},
			},
		},
		2,
	},
	{
		Tree{
			0: map[int]struct{}{
				2: {},
				4: {},
				1: {},
			},
			2: map[int]struct{}{
				3: {},
			},
			4: map[int]struct{}{
				5: {},
			},
			5: map[int]struct{}{
				6: {},
				7: {},
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

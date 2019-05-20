package day270

import "testing"

// nolint
var testcases = []struct {
	n          int
	edges      []Edge
	timeForAll int
}{
	{5, []Edge{{0, 1, 5}, {0, 2, 3}, {0, 5, 4}, {1, 3, 8}, {2, 3, 1}, {3, 5, 10}, {3, 4, 5}}, 9},
}

func TestTimeAllNodesReceiveMessage(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := TimeAllNodesReceiveMessage(tc.n, tc.edges); result != tc.timeForAll {
			t.Errorf("Expected %v, got %v", tc.timeForAll, result)
		}
	}
}

func BenchmarkTimeAllNodesReceiveMessage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			TimeAllNodesReceiveMessage(tc.n, tc.edges)
		}
	}
}

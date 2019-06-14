package day294

import "testing"

// nolint
var testcases = []struct {
	elevations    Elevation
	paths         PathDistance
	shortestRoute int
	err           error
}{
	{
		Elevation{
			0: 15,
			1: 10,
			2: 25,
		},
		PathDistance{
			0: {
				2: 5,
			},
			1: {
				0: 1,
			},
			2: {
				1: 7,
			},
		},
		0,
		ErrNotPossible(),
	},
	{
		Elevation{
			0: 10,
			1: 25,
			2: 15,
		},
		PathDistance{
			0: {
				1: 15,
				2: 5,
			},
			1: {
				0: 1,
			},
			2: {
				1: 7,
				0: 30,
			},
		},
		13,
		nil,
	},
	{
		Elevation{
			0: 10,
			1: 5,
			2: 7,
		},
		PathDistance{
			0: {
				1: 5,
				2: 10,
			},
			1: {
				0: 5,
			},
			2: {
				0: 10,
			},
		},
		0,
		ErrNotPossible(),
	},
	{
		Elevation{
			0: 10,
			1: 5,
			2: 15,
		},
		PathDistance{
			0: {
				1: 5,
				2: 10,
			},
			1: {
				0: 5,
			},
			2: {
				0: 10,
			},
		},
		20,
		nil,
	},
	{
		Elevation{
			0: 5,
			1: 25,
			2: 15,
			3: 20,
			4: 10,
		},
		PathDistance{
			0: {
				1: 10,
				2: 8,
				3: 15,
			},
			1: {
				3: 12,
			},
			2: {
				4: 10,
			},
			3: {
				4: 5,
				0: 17,
			},
			4: {
				0: 10,
			},
		},
		28,
		nil,
	},
}

func TestShortestRunnerPath(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result, err := ShortestRunnerPath(tc.elevations, tc.paths); result != tc.shortestRoute || err != tc.err {
			t.Errorf("Expected (%v,%v), got (%v,%v)", tc.shortestRoute, tc.err, result, err)
		}
	}
}

func BenchmarkShortestRunnerPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ShortestRunnerPath(tc.elevations, tc.paths) // nolint
		}
	}
}

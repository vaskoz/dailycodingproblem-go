package day328

import (
	"math"
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	scores    []float64
	kFactor   float64
	matches   []Match
	expected  []float64
	tolerance float64
}{
	{
		[]float64{1000, 1000, 1000, 1000, 1000},
		30,
		[]Match{{0, 1}, {0, 2}, {0, 3}, {0, 4}},
		[]float64{1056.24, 985, 985.65, 986.26, 986.85},
		100,
	},
}

func TestEloGames(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		result := EloGames(tc.scores, tc.kFactor, tc.matches)
		for i := range result {
			result[i] = math.Round(result[i]*tc.tolerance) / tc.tolerance
		}

		if !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkEloGames(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			EloGames(tc.scores, tc.kFactor, tc.matches)
		}
	}
}

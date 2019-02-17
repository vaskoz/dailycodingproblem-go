package day175

import "testing"

var testcases = []struct {
	start             State
	states            Transition
	numSteps          int
	expected          map[State]int
	expectedTolerance int
}{
	{"a", Transition{
		"a": {
			"a": 0.9,
			"b": 0.075,
			"c": 0.025,
		},
		"b": {
			"a": 0.15,
			"b": 0.8,
			"c": 0.05,
		},
		"c": {
			"a": 0.25,
			"b": 0.25,
			"c": 0.5,
		},
	}, 5000, map[State]int{
		"a": 3100,
		"b": 1500,
		"c": 300,
	}, 500},
}

func TestRunMarkovChain(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := RunMarkovChain(tc.start, tc.states, tc.numSteps); !close(result, tc.expected, tc.expectedTolerance) {
			t.Errorf("Expected %v, got %v", result, tc.expected)
		}
	}
}

func BenchmarkRunMarkovChain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			RunMarkovChain(tc.start, tc.states, tc.numSteps)
		}
	}
}

func close(a, b map[State]int, tolerance int) bool {
	for s, cnt := range a {
		if abs(cnt-b[s]) > tolerance {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

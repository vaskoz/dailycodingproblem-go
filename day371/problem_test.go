package day371

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	equations   string
	expected    map[string]int
	expectedErr error
}{
	{
		`y = x + 1
	5 = x + 3
	10 = z + y + 2`,
		map[string]int{
			"x": 2,
			"y": 3,
			"z": 5,
		},
		nil,
	},
	{
		`y = x + 1
	10 = z + y + 2`,
		nil,
		errNoAnswerPossible,
	},
	{
		`
		foo+1+foo+1+foo = foo + 12
		 1 = bar+foo+2`,
		map[string]int{
			"foo": 5,
			"bar": -6,
		},
		nil,
	},
}

func TestSolveArithmeticEquations(t *testing.T) {
	for _, tc := range testcases {
		if result, err := SolveArithmeticEquations(tc.equations); !reflect.DeepEqual(result, tc.expected) ||
			err != tc.expectedErr {
			t.Errorf("Expected (%v,%v), got (%v,%v)", tc.expected, tc.expectedErr, result, err)
		}
	}
}

func TestSolveArithmeticEquationsBadInput(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic when bad input")
		}
	}()
	// nolint
	SolveArithmeticEquations(`a+y=1 x+z=2
x=5`)
}

func BenchmarkSolveArithmeticEquations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SolveArithmeticEquations(tc.equations) // nolint
		}
	}
}

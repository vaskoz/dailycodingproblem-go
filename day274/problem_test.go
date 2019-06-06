package day274

import "testing"

// nolint
var testcases = []struct {
	expr     string
	expected int
}{
	{"-1 + (2 + 3)", 4},
	{"-2 + (2 + (3 + 5 + 8))", 16},
	{"-2 + (-2 + (-3 + 5 + 8))", 6},
	{"2 - (2 - (3 - 5 - 8))", -10},
}

func TestEvalSimpleMathExpression(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := EvalSimpleMathExpression(tc.expr); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkEvalSimpleMathExpression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			EvalSimpleMathExpression(tc.expr)
		}
	}
}

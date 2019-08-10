package day352

import "testing"

// nolint
var testcases = []struct {
	puzzle  [][]Square
	isValid bool
}{
	{
		[][]Square{
			{W, W, W, W, B, W, W, W, W, W, B, W, W, W, W},
			{W, W, W, W, B, W, W, W, W, W, B, W, W, W, W},
			{W, W, W, W, W, W, W, W, W, W, W, W, W, W, W},
			{W, W, W, B, W, W, W, B, W, W, W, W, W, W, W},
			{W, W, W, W, W, W, B, W, W, W, W, W, B, B, B},
			{B, B, W, W, W, B, W, W, W, B, W, W, W, W, W},
			{W, W, W, W, B, W, W, W, W, W, W, W, W, W, W},
			{W, W, W, B, W, W, W, W, W, W, W, B, W, W, W},
			{W, W, W, W, W, W, W, W, W, W, B, W, W, W, W},
			{W, W, W, W, W, B, W, W, W, B, W, W, W, B, B},
			{B, B, B, W, W, W, W, W, B, W, W, W, W, W, W},
			{W, W, W, W, W, W, W, B, W, W, W, B, W, W, W},
			{W, W, W, W, W, W, W, W, W, W, W, W, W, W, W},
			{W, W, W, W, B, W, W, W, W, W, B, W, W, W, W},
			{W, W, W, W, B, W, W, W, W, W, B, W, W, W, W},
		},
		true,
	},
	{
		[][]Square{
			{W, W, W, W, B, W, W, W, W, W, B, W, W, W, W},
			{W, W, W, W, B, W, W, W, W, W, B, W, W, W, W},
			{W, W, W, W, W, W, W, W, W, W, W, W, W, W, W},
			{W, W, W, B, W, W, W, B, W, W, W, W, W, W, W},
			{W, W, W, W, W, W, B, W, W, W, W, W, B, B, B},
			{B, B, W, W, W, B, W, W, W, B, W, W, W, W, W},
			{W, W, W, W, B, W, W, W, W, W, W, W, W, W, W},
			{W, W, W, B, W, W, W, W, W, W, W, B, W, W, W},
			{W, W, W, W, W, W, W, W, W, W, B, W, W, W, W},
			{W, W, W, W, W, B, W, W, W, B, W, W, W, B, B},
			{B, B, B, W, W, W, W, W, B, W, W, W, W, W, W},
			{W, W, W, W, W, W, W, B, W, W, W, B, W, W, W},
			{W, W, W, W, W, W, W, W, W, W, W, W, W, W, W},
			{W, W, W, W, B, W, W, W, W, W, B, W, W, W, W},
			{W, W, B, W, B, W, W, W, W, W, B, W, W, W, W},
		},
		false,
	},
	{
		[][]Square{
			{W, W, W, W, B, W, W, W, W, W, B, W, W, W, W},
			{W, W, W, W, B, W, W, W, W, W, B, W, W, W, W},
			{W, W, W, W, W, W, W, W, W, W, W, W, W, W, W},
			{W, W, W, B, W, W, W, B, W, W, W, W, W, W, W},
			{W, W, W, W, W, W, B, W, W, W, W, W, B, B, B},
			{B, B, W, W, W, B, W, W, W, B, W, W, W, W, W},
			{W, W, W, W, B, W, W, W, W, W, W, W, W, W, W},
			{W, W, W, B, W, W, W, W, W, W, W, B, W, W, W},
			{W, W, W, W, W, W, W, W, W, W, B, W, W, W, W},
			{W, W, W, W, W, B, W, W, W, B, W, W, W, B, B},
			{B, B, B, W, W, W, W, W, B, W, W, W, W, W, W},
			{W, W, W, W, W, W, W, B, W, W, W, B, W, W, W},
			{W, W, W, W, W, W, W, W, W, B, W, W, W, W, W},
			{W, W, W, W, B, W, W, W, W, W, B, W, W, W, W},
			{W, W, W, W, B, W, W, W, W, W, B, W, W, W, W},
		},
		false,
	},
}

func TestIsValidCrossword(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		if valid := IsValidCrossword(tc.puzzle); valid != tc.isValid {
			t.Errorf("TCID %d. Expected %v, got %v", tcid, tc.isValid, valid)
		}
	}
}

func BenchmarkIsValidCrossword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			IsValidCrossword(tc.puzzle)
		}
	}
}

package day85

import "testing"

var testcases = []struct {
	x, y, b, expected int32
}{
	{10, 20, 1, 10},
	{10, 20, 0, 20},
	{-20, -40, 1, -20},
	{-20, -40, 0, -40},
	{100, -200, 1, 100},
	{100, -200, 0, -200},
}

func TestMathematicalIf(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MathematicalIf(tc.x, tc.y, tc.b); tc.expected != result {
			t.Errorf("x(%d),y(%d),b(%d),expected(%d),got(%d)", tc.x, tc.y, tc.b, tc.expected, result)
		}
	}
}

func BenchmarkMathematicalIf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MathematicalIf(tc.x, tc.y, tc.b)
		}
	}
}

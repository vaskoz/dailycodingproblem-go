package day50

import "testing"

var testcases = []struct {
	tree     *ArithmeticTree
	expected int
}{
	{&ArithmeticTree{op: MUL,
		left:  &ArithmeticTree{op: ADD, left: &ArithmeticTree{val: 3}, right: &ArithmeticTree{val: 2}},
		right: &ArithmeticTree{op: ADD, left: &ArithmeticTree{val: 4}, right: &ArithmeticTree{val: 5}}},
		45},
	{&ArithmeticTree{op: SUB,
		left:  &ArithmeticTree{op: DIV, left: &ArithmeticTree{val: 100}, right: &ArithmeticTree{val: 2}},
		right: &ArithmeticTree{op: DIV, left: &ArithmeticTree{val: 100}, right: &ArithmeticTree{val: 5}}},
		30},
}

func TestCalculate(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := Calculate(tc.tree); result != tc.expected {
			t.Errorf("Expected %d got %d", tc.expected, result)
		}
	}
}

func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			Calculate(tc.tree)
		}
	}
}

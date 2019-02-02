package day163

import "testing"

var testcases = []struct {
	input    []string
	expected int
}{
	{[]string{"5", "3", "+"}, 8},
	{[]string{"15", "7", "1", "1", "+", "-", "/", "3", "*", "2", "1", "1", "+", "+", "-"}, 5},
}

func TestReversePolishCalculator(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := ReversePolishCalculator(tc.input); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func TestReversePolishCalculatorBadInput(t *testing.T) {
	t.Parallel()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expect a panic for bad input")
		}
	}()
	ReversePolishCalculator([]string{"15", "20", "abc", "?"})
}

func TestReversePolishCalculatorBadInput2(t *testing.T) {
	t.Parallel()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expect a panic for bad input")
		}
	}()
	ReversePolishCalculator([]string{"15", "20", "-", "11"})
}

func BenchmarkReversePolishCalculator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ReversePolishCalculator(tc.input)
		}
	}
}

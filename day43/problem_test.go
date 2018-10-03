package day43

import "testing"

var testcases = []struct {
	input []int
	max   []int
}{
	{[]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
}

func TestMaxIntStack(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		stack := NewMaxIntStack()
		for i, push := range tc.input {
			stack.Push(push)
			if max, err := stack.Max(); err != nil || max != tc.max[i] {
				t.Errorf("TC%d push expected,got max (%d,%d) error should be nil (%v)",
					tcid, tc.max[i], max, err)
			}
		}
		for i := range tc.input {
			if max, err := stack.Max(); err != nil || max != tc.max[len(tc.max)-1-i] {
				t.Errorf("TC%d pre-pop expected,got max (%d,%d) error should be nil (%v)",
					tcid, tc.max[len(tc.max)-1-i], max, err)
			}
			if val, err := stack.Pop(); err != nil || val != tc.input[len(tc.input)-1-i] {
				t.Errorf("TC%d pop expected,got val (%d,%d) error should be nil (%v)",
					tcid, tc.input[len(tc.input)-1-i], val, err)
			}
		}
		if _, err := stack.Pop(); err != StackEmptyError() {
			t.Errorf("Stack should be empty and return an error on Pop()")
		}
		if _, err := stack.Max(); err != StackEmptyError() {
			t.Errorf("Stack should be empty and return an error on Max()")
		}
	}
}

func BenchmarkMaxIntStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			stack := NewMaxIntStack()
			for _, push := range tc.input {
				stack.Push(push)
				stack.Max()
				stack.Pop()
			}
		}
	}
}

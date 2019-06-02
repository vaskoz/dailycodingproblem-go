package day283

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	n        int
	expected []int
}{
	{0, []int{}},
	{1, []int{1}},
	{2, []int{1, 2}},
	{11, []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15}},
}

func TestRegularNumbersBrute(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := RegularNumbersBrute(tc.n); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func TestRegularNumbersBruteBadInput(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic for a negative input")
		}
	}()
	RegularNumbersBrute(-2)
}

func BenchmarkRegularNumbersBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			RegularNumbersBrute(tc.n)
		}
	}
}

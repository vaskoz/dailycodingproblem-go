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
	{26, []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15, 16, 18, 20, 24, 25, 27, 30, 32, 36, 40, 45, 48, 50, 54, 60}},
}

func TestRegularNumbersFaster(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := RegularNumbersFaster(tc.n); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func TestRegularNumbersFasterBadInput(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic for a negative input")
		}
	}()
	RegularNumbersFaster(-2)
}

func BenchmarkRegularNumbersFaster(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			RegularNumbersFaster(tc.n)
		}
	}
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

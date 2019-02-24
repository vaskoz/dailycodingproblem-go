package day181

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	input    string
	expected []string
}{
	{"racecarannakayak", []string{"racecar", "anna", "kayak"}},
	{"abc", []string{"a", "b", "c"}},
	{"racecarracecar", []string{"racecarracecar"}},
	{"a", []string{"a"}},
	{"", nil},
}

func TestMinimumPartitionPalindrome(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MinimumPartitionPalindrome(tc.input); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkMinimumPartitionPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MinimumPartitionPalindrome(tc.input)
		}
	}
}

package day57

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	long     string
	k        int
	expected []string
}{
	{"the quick brown fox jumps over the lazy dog",
		10,
		[]string{"the quick", "brown fox", "jumps over", "the lazy", "dog"},
	},
	{"the quick brown fox jumps over the lazy dog",
		4,
		nil,
	},
}

func TestBreakByLength(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := BreakByLength(tc.long, tc.k); !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

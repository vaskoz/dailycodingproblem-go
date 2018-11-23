package day92

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	prereqs  map[string][]string
	expected []string
}{
	{map[string][]string{
		"CSC300": []string{"CSC100", "CSC200"},
		"CSC200": []string{"CSC100"},
		"CSC100": []string{},
	},
		[]string{"CSC100", "CSC200", "CSC300"}},
}

func TestCourseOrder(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := CourseOrder(tc.prereqs); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkCourseOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CourseOrder(tc.prereqs)
		}
	}
}

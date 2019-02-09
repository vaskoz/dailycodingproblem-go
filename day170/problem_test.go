package day170

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	start, end  string
	dict        []string
	expected    []string
	expectedErr error
}{
	{"dog", "cat", []string{"dot", "dop", "dat", "cat"}, []string{"dog", "dot", "dat", "cat"}, nil},
	{"dog", "cat", []string{"dot", "tod", "dat", "dar"}, nil, ErrTransformationNotPossible()},
}

func TestShortestTransformation(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result, err := ShortestTransformation(tc.start, tc.end, tc.dict); err != tc.expectedErr || !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected (%v,%v) got (%v,%v)", tc.expected, tc.expectedErr, result, err)
		}
	}
}

func BenchmarkShortestTransformation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ShortestTransformation(tc.start, tc.end, tc.dict)
		}
	}
}

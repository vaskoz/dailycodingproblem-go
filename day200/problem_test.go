package day200

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	X []Interval
	P []Point
}{
	{[]Interval{{1, 4}, {4, 5}, {7, 9}, {9, 12}}, []Point{4, 9}},
	{[]Interval{}, nil},
	{nil, nil},
	{[]Interval{{9, 12}, {7, 9}, {4, 5}, {1, 4}}, []Point{4, 9}},
	{[]Interval{{9, 12}, {7, 9}, {4, 5}, {1, 4}, {3, 5}}, []Point{4, 9}},
}

func TestStabIntervalPoints(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		input := append([]Interval{}, tc.X...)
		if result := StabIntervalPoints(input); !reflect.DeepEqual(result, tc.P) {
			t.Errorf("Expected %v, got %v", tc.P, result)
		}
	}
}

func BenchmarkStabIntervalPoints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			input := append([]Interval{}, tc.X...)
			StabIntervalPoints(input)
		}
	}
}

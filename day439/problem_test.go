package day439

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	segments []Flight
	start    []Airport
	expected []Airport
}{
	{[]Flight{{"HNL", "AKL"}, {"YUL", "ORD"}, {"ORD", "SFO"}, {"SFO", "HNL"}},
		[]Airport{"YUL"},
		[]Airport{"YUL", "ORD", "SFO", "HNL", "AKL"}},
	{[]Flight{{"HNL", "AKL"}, {"YUL", "ORD"}, {"ORD", "SFO"}, {"SFO", "HNL"}, {"ORD", "LAX"}},
		[]Airport{"YUL"},
		nil},
	{[]Flight{{"SFO", "HKO"}, {"YYZ", "SFO"}, {"YUL", "YYZ"}, {"HKO", "ORD"}},
		[]Airport{"YUL"},
		[]Airport{"YUL", "YYZ", "SFO", "HKO", "ORD"}},
	{[]Flight{{"SFO", "COM"}, {"COM", "YYZ"}},
		[]Airport{"COM"},
		nil},
	{[]Flight{{"A", "B"}, {"A", "C"}, {"B", "C"}, {"C", "A"}},
		[]Airport{"A"},
		[]Airport{"A", "B", "C", "A", "C"}},
}

func TestItinerary(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := Itinerary(tc.segments, tc.start); !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v but got %v", tc.expected, result)
		}
	}
}

func BenchmarkItinerary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			Itinerary(tc.segments, tc.start)
		}
	}
}

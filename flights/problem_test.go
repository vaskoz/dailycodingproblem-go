package flights

import (
	"reflect"
	"testing"
)

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

package day346

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	flights             Flights
	origin, destination string
	connections         int
	itinerary           []string
	cost                int
}{
	{
		Flights{
			"JFK": {
				"ATL": 150,
				"HKG": 800,
				"LAX": 500,
			},
			"ATL": {
				"SFO": 400,
				"JFK": 1,
				"ORD": 90,
			},
			"ORD": {
				"LAX": 200,
			},
			"LAX": {
				"DFW": 80,
			},
		},
		"JFK",
		"LAX",
		3,
		[]string{"JFK", "ATL", "ORD", "LAX"},
		440,
	},
	{
		Flights{
			"JFK": {
				"ATL": 150,
				"HKG": 800,
				"LAX": 500,
			},
			"ATL": {
				"SFO": 400,
				"ORD": 90,
			},
			"ORD": {
				"LAX": 200,
			},
			"LAX": {
				"DFW": 80,
			},
		},
		"JFK",
		"LAX",
		1,
		[]string{"JFK", "LAX"},
		500,
	},
}

func TestCheapestItinerary(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if itinerary, price := CheapestItinerary(tc.flights, tc.origin,
			tc.destination, tc.connections); !reflect.DeepEqual(tc.itinerary, itinerary) ||
			tc.cost != price {
			t.Errorf("Expected (%v,%v), got (%v,%v)", tc.itinerary, tc.cost, itinerary, price)
		}
	}
}

func BenchmarkCheapestItinerary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CheapestItinerary(tc.flights, tc.origin, tc.destination, tc.connections)
		}
	}
}

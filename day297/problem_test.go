package day297

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	prefs           map[Customer][]Drink
	memorizedDrinks []Drink
}{
	{
		map[Customer][]Drink{
			0: {0, 1, 3, 6},
			1: {1, 4, 7},
			2: {2, 4, 7, 5},
			3: {3, 2, 5},
			4: {5, 8},
		},
		[]Drink{1, 5},
	},
}

func TestLazyBartender(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := LazyBartender(tc.prefs); !reflect.DeepEqual(result, tc.memorizedDrinks) {
			t.Errorf("Expected %v, got %v", tc.memorizedDrinks, result)
		}
	}
}

func BenchmarkLazyBartender(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LazyBartender(tc.prefs)
		}
	}
}

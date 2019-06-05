package day286

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	buildings []Building
	skyline   []Skyline
}{
	{
		[]Building{{0, 15, 3}, {4, 11, 5}, {19, 23, 4}},
		[]Skyline{{0, 3}, {4, 5}, {11, 3}, {15, 0}, {19, 4}, {23, 0}},
	},
}

func TestCreateSkyline(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if skyline := CreateSkyline(tc.buildings); !reflect.DeepEqual(skyline, tc.skyline) {
			t.Errorf("Expected %v, got %v", tc.skyline, skyline)
		}
	}
}

func BenchmarkCreateSkyline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CreateSkyline(tc.buildings)
		}
	}
}

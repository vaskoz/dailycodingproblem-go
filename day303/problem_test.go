package day303

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	hh, mm int
	angle  int
}{
	{12, 00, 0},
	{11, 00, 30},
}

func TestAngleClockHands(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if angle := AngleClockHands(tc.hh, tc.mm); angle != tc.angle {
			t.Errorf("Expected %v, got %v", tc.angle, angle)
		}
	}
}

func TestAngleClockHandsBadTime(t *testing.T) {
	t.Parallel()
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic with an invalid time")
		}
	}()
	AngleClockHands(20, 60)
}

func BenchmarkAngleClockHands(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			AngleClockHands(tc.hh, tc.mm)
		}
	}
}

func TestTimesOverlappingHands(t *testing.T) {
	t.Parallel()
	result := TimesOverlappingHands()
	expected := [][]int{
		{0, 0}, {1, 5}, {2, 10}, {3, 15}, {4, 20}, {5, 25}, {6, 30}, {7, 35}, {8, 40}, {9, 45}, {10, 50}, {11, 55},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", result, expected)
	}
}

func BenchmarkTimesOverlappingHands(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TimesOverlappingHands()
	}
}

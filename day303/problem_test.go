package day303

import "testing"

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

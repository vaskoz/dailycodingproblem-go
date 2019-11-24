package day369

import (
	"math"
	"testing"
)

// nolint
var testcases = []struct {
	action        string
	data          Datapoint
	min, max, avg Price
}{
	{"addOrUpdate", Datapoint{10, 100.0}, 100.0, 100.0, 100.0},
	{"addOrUpdate", Datapoint{20, 500.0}, 100.0, 500.0, 300.0},
	{"addOrUpdate", Datapoint{10, 1000.0}, 500.0, 1000.0, 750.0},
	{"remove", Datapoint{20, 500.0}, 1000.0, 1000.0, 1000.0},
	{"remove", Datapoint{10, 1000.0}, math.MaxFloat64, -math.MaxFloat64, 0},
	{"addOrUpdate", Datapoint{100, 10000.0}, 10000.0, 10000.0, 10000.0},
	{"addOrUpdate", Datapoint{200, 10000.0}, 10000.0, 10000.0, 10000.0},
}

func TestStockService(t *testing.T) {
	t.Parallel()

	ss := NewStockService()

	for _, tc := range testcases {
		switch tc.action {
		case "addOrUpdate":
			ss.AddOrUpdate(tc.data)
		case "remove":
			ss.Remove(tc.data.Timestamp)
		}

		if min := ss.Min(); min != tc.min {
			t.Errorf("Expected min %v, got %v", tc.min, min)
		}

		if max := ss.Max(); max != tc.max {
			t.Errorf("Expected max %v, got %v", tc.max, max)
		}

		if avg := ss.Avg(); avg != tc.avg {
			t.Errorf("Expected avg %v, got %v", tc.avg, avg)
		}
	}
}

func BenchmarkStockService(b *testing.B) {
	ss := NewStockService()

	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			switch tc.action {
			case "addOrUpdate":
				ss.AddOrUpdate(tc.data)
			case "remove":
				ss.Remove(tc.data.Timestamp)
			}

			ss.Min()
			ss.Max()
			ss.Avg()
		}
	}
}

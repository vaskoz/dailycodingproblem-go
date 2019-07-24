package day325

import (
	"testing"
)

// nolint
var testcases = []struct {
	val       float64
	from, to  Unit
	expected  float64
	err       error
	tolerance float64
}{
	{
		10.0,
		"in", "cm",
		25.4,
		nil,
		0.05,
	},
	{
		10.0,
		"in", "km",
		0.000254,
		nil,
		0.0000005,
	},
	{
		10.0,
		"in", "km",
		0.000254,
		nil,
		0.0000005,
	},
	{
		10.0,
		"in", "ml",
		0.0,
		ErrConversionImpossible(),
		0.0,
	},
}

func TestConverter(t *testing.T) {
	t.Parallel()
	c := NewConverter()
	c.AddConversion(12.0, "ft", "in")
	c.AddConversion(5280.0, "mi", "ft")
	c.AddConversion(1.609, "mi", "km")
	c.AddConversion(1000, "km", "m")
	c.AddConversion(100, "m", "cm")
	c.AddConversion(3.0, "yd", "ft")
	c.AddConversion(22.0, "chain", "yd")
	for _, tc := range testcases {
		if result, err := c.Convert(tc.val, tc.from, tc.to); err != tc.err || abs(result-tc.expected) > tc.tolerance {
			t.Errorf("Expected (%v,%v), got (%v,%v)", tc.expected, tc.err, result, err)
		}
	}
}

func BenchmarkConverter(b *testing.B) {
	c := NewConverter()
	c.AddConversion(12.0, "in", "ft")
	c.AddConversion(5280.0, "ft", "mi")
	c.AddConversion(1.609, "mi", "km")
	c.AddConversion(0.001, "km", "m")
	c.AddConversion(0.01, "m", "cm")
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			c.Convert(tc.val, tc.from, tc.to) // nolint
		}
	}
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

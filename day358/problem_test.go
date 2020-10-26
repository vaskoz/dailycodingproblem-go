package day358

import "testing"

// nolint
var actions = []struct {
	key    interface{}
	action string
	count  int
}{
	{nil, "max", 0},
	{nil, "min", 0},
	{"cat", "plus", 1},
	{"dog", "plus", 1},
	{"cat", "minus", 1},
	{"dog", "minus", 1},
	{"foo", "plus", 3},
	{"bar", "plus", 4},
	{"bar", "max", 4},
	{"foo", "min", 3},
	{"foo", "plus", 2},
	{"foo", "max", 5},
	{"bar", "min", 4},
	{"foo", "minus", 3},
	{"bar", "max", 4},
	{"foo", "min", 2},
	{"foo", "minus", 1},
	{"bar", "max", 4},
	{"foo", "min", 1},
	{"foo", "minus", 1},
	{"bar", "max", 4},
	{"bar", "min", 4},
	{"doesntexist", "minus", 4},
}

func TestBigOhOne(t *testing.T) {
	t.Parallel()

	bo := New()

	for _, a := range actions {
		switch a.action {
		case "plus":
			for i := 0; i < a.count; i++ {
				bo.Plus(a.key)
			}
		case "minus":
			for i := 0; i < a.count; i++ {
				bo.Minus(a.key)
			}
		case "min":
			if key, count := bo.Min(); key != a.key || count != a.count {
				t.Errorf("Expected Min (%v,%v) got (%v,%v)", a.key, a.count, key, count)
			}
		case "max":
			if key, count := bo.Max(); key != a.key || count != a.count {
				t.Errorf("Expected Max (%v,%v) got (%v,%v)", a.key, a.count, key, count)
			}
		}
	}
}

func BenchmarkBigOhOne(b *testing.B) {
	bo := New()

	for i := 0; i < b.N; i++ {
		for _, a := range actions {
			switch a.action {
			case "plus":
				for i := 0; i < a.count; i++ {
					bo.Plus(a.key)
				}
			case "minus":
				for i := 0; i < a.count; i++ {
					bo.Minus(a.key)
				}
			case "min":
				bo.Min()
			case "max":
				bo.Max()
			}
		}
	}
}

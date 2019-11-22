package day367

import "testing"

// nolint
var (
	intLessThanEqual LessThanEqual = func(a, b interface{}) bool {
		aI := a.(int)
		bI := b.(int)
		return aI <= bI
	}
	stringLessThanEqual LessThanEqual = func(a, b interface{}) bool {
		aS := a.(string)
		bS := b.(string)
		return aS <= bS
	}
)

// nolint
var testcases = []struct {
	a                 []interface{}
	aLessThanEqual    LessThanEqual
	b                 []interface{}
	bLessThanEqual    LessThanEqual
	combLessThanEqual LessThanEqual

	expected []interface{}
}{
	{
		[]interface{}{5, 10, 15},
		intLessThanEqual,
		[]interface{}{3, 8, 9},
		intLessThanEqual,
		intLessThanEqual,
		[]interface{}{3, 5, 8, 9, 10, 15},
	},
	{
		[]interface{}{3, 8, 9},
		intLessThanEqual,
		[]interface{}{5, 10, 15},
		intLessThanEqual,
		intLessThanEqual,
		[]interface{}{3, 5, 8, 9, 10, 15},
	},
	{
		[]interface{}{9, 3, 8, 1, 9},
		intLessThanEqual,
		[]interface{}{25, 2, 100, 10, 15},
		intLessThanEqual,
		intLessThanEqual,
		[]interface{}{1, 2, 3, 8, 9, 9, 10, 15, 25, 100},
	},
	{
		[]interface{}{"o", "a", "k", "s", "v"},
		stringLessThanEqual,
		[]interface{}{"w", "e", "l", "c", "o", "m", "e"},
		stringLessThanEqual,
		stringLessThanEqual,
		[]interface{}{"a", "c", "e", "e", "k", "l", "m", "o", "o", "s", "v", "w"},
	},
	{
		[]interface{}{10, 8, 9, 7, 5, 6, 4, 2, 3, 1, 0},
		intLessThanEqual,
		[]interface{}{"a", "c", "k", "b", "d", "f", "e", "h", "g", "i", "j"},
		stringLessThanEqual,
		func(a, b interface{}) bool {
			if av, ok := a.(string); ok {
				if bv, ok := b.(string); ok {
					return stringLessThanEqual(av, bv)
				} else {
					avr := int(av[0]) - 'a'
					bv := b.(int)
					return avr <= bv
				}
			} else {
				av := a.(int)
				if bv, ok := b.(string); ok {
					bvr := int(bv[0]) - 'a'
					return av <= bvr
				} else {
					bv := b.(int)
					return intLessThanEqual(av, bv)
				}
			}
		},
		[]interface{}{0, "a", 1, "b", 2, "c", 3, "d", 4, "e", 5, "f", 6, "g", 7, "h",
			8, "i", 9, "j", 10, "k"},
	},
}

func TestMergeSortedIterators(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		a := NewSortedIterator(tc.a, tc.aLessThanEqual)
		b := NewSortedIterator(tc.b, tc.bLessThanEqual)
		c := MergeSortedIterators(a, b, tc.combLessThanEqual)

		for index, v := range tc.expected {
			if peek := c.Peek(); v != peek {
				t.Errorf("Expected element index %d equal to %v, got %v", index, v, peek)
			} else if got := c.Next(); peek != got {
				t.Errorf("peek and got do not match index %d peek=%v got=%v", index, peek, got)
			}
		}

		if !c.Empty() {
			t.Errorf("Expected an empty merged iterator, but it's not")
		}

		if c.Peek() != nil || c.Next() != nil {
			t.Errorf("once empty, all peek and next should be nil")
		}
	}
}

func BenchmarkMergeSortedIterators(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			a := NewSortedIterator(tc.a, tc.aLessThanEqual)
			b := NewSortedIterator(tc.b, tc.bLessThanEqual)
			c := MergeSortedIterators(a, b, tc.combLessThanEqual)

			for !c.Empty() {
				c.Next()
			}
		}
	}
}

func TestSingleSortedIterators(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		a := NewSortedIterator(tc.a, tc.aLessThanEqual)

		first := a.Next()

		for !a.Empty() {
			next := a.Next()
			if !tc.aLessThanEqual(first, next) {
				t.Errorf("a not sorted %v should be less than %v", first, next)
			}

			first = next
		}

		if a.Peek() != nil || a.Next() != nil || !a.Empty() {
			t.Errorf("expected empty with nils")
		}

		b := NewSortedIterator(tc.b, tc.bLessThanEqual)

		first = b.Next()

		for !b.Empty() {
			next := b.Next()
			if !tc.bLessThanEqual(first, next) {
				t.Errorf("b not sorted")
			}

			first = next
		}

		if a.Peek() != nil || a.Next() != nil || !a.Empty() {
			t.Errorf("expected empty with nils")
		}
	}
}

func BenchmarkSingleSortedIterators(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			a := NewSortedIterator(tc.a, tc.aLessThanEqual)
			b := NewSortedIterator(tc.b, tc.bLessThanEqual)

			for !a.Empty() {
				a.Next()
			}

			for !b.Empty() {
				b.Next()
			}
		}
	}
}

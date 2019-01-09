package day139

import "testing"

var testcases = []struct {
	data []interface{}
}{
	{[]interface{}{3, 5, 6, 7, 9}},
	{[]interface{}{"foo", "bar", "baz"}},
}

type SliceIterator struct {
	Slice []interface{}
	pos   int
}

func (si *SliceIterator) Next() interface{} {
	if si.pos >= len(si.Slice) {
		return nil
	}
	result := si.Slice[si.pos]
	si.pos++
	return result
}

func (si *SliceIterator) HasNext() bool {
	return si.pos < len(si.Slice)
}

func TestPeekableIterator(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		p := PeekableIterator{Iter: &SliceIterator{Slice: tc.data}}
		for i := range tc.data {
			if result := p.Peek(); result != tc.data[i] {
				t.Errorf("Peek should return %v got %v", tc.data[i], result)
			}
			if more := p.HasNext(); !more {
				t.Errorf("There should be more data")
			}
			if result := p.Next(); result != tc.data[i] {
				t.Errorf("Next should return %v got %v", tc.data[i], result)
			}
		}
		if result := p.Peek(); result != nil {
			t.Errorf("Peek should return %v got %v", nil, result)
		}
		if more := p.HasNext(); more {
			t.Errorf("There should be NO more data")
		}
		if result := p.Next(); result != nil {
			t.Errorf("Next should return %v got %v", nil, result)
		}
	}
}

func TestPeekableIteratorNoPeek(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		p := PeekableIterator{Iter: &SliceIterator{Slice: tc.data}}
		for i := range tc.data {
			if more := p.HasNext(); !more {
				t.Errorf("There should be more data")
			}
			if result := p.Next(); result != tc.data[i] {
				t.Errorf("Next should return %v got %v", tc.data[i], result)
			}
		}
		if more := p.HasNext(); more {
			t.Errorf("There should be NO more data")
		}
		if result := p.Next(); result != nil {
			t.Errorf("Next should return %v got %v", nil, result)
		}
	}
}

func BenchmarkPeekableIterator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			p := PeekableIterator{Iter: &SliceIterator{Slice: tc.data}}
			for range tc.data {
				p.Peek()
				p.HasNext()
				p.Next()
			}
		}
	}
}

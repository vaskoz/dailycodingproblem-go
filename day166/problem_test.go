package day166

import "testing"

var testcases = []struct {
	input    [][]interface{}
	expected []interface{}
}{
	{[][]interface{}{
		{1, 2}, {3}, {}, {4, 5, 6},
	},
		[]interface{}{1, 2, 3, 4, 5, 6},
	},
	{nil, []interface{}{}},
	{[][]interface{}{{}, {}, {}}, []interface{}{}},
}

func TestTwoDimensionalIterator(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		tdi := NewTwoDimensionalIterator(tc.input)
		for _, v := range tc.expected {
			if !tdi.HasNext() {
				t.Error("HasNext should return true")
			}
			if result, err := tdi.Next(); result != v || err != nil {
				t.Errorf("Expected (%v,nil) and got (%v,%v)", v, result, err)
			}
		}
		if tdi.HasNext() {
			t.Error("HasNext should return false")
		}
		if _, err := tdi.Next(); err == nil {
			t.Errorf("Expected an error on Next when HasNext is false")
		}
	}
}

func BenchmarkTwoDimensionalIterator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			tdi := NewTwoDimensionalIterator(tc.input)
			for range tc.expected {
				tdi.Next()
				tdi.HasNext()
			}
		}
	}
}

package day180

import "testing"

var testcases = []struct {
	input, expected []interface{}
}{
	{[]interface{}{1, 2, 3, 4, 5}, []interface{}{1, 5, 2, 4, 3}},
	{[]interface{}{1, 2, 3, 4}, []interface{}{1, 4, 2, 3}},
	{[]interface{}{}, []interface{}{}},
	{nil, nil},
}

func TestInterleaveStack(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		var s Stack
		for _, v := range tc.input {
			s.Push(v)
		}
		InterleaveStack(s)
		for i := range tc.expected {
			v := tc.expected[len(tc.expected)-1-i]
			if result := s.Pop(); result != v {
				t.Errorf("Expected (%v) got (%v)", v, result)
			}
		}
	}
}

func BenchmarkInterleaveStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			var s Stack
			for _, v := range tc.input {
				s.Push(v)
			}
			InterleaveStack(s)
		}
	}
}

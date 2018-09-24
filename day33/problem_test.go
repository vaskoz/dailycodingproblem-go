package day33

import (
	"container/heap"
	"testing"
)

var testcases = []struct {
	input    []int
	expected []float64
}{
	{[]int{2, 1, 5, 7, 2, 0, 5}, []float64{2, 1.5, 2, 3.5, 2, 2, 2}},
}

func TestRunningMedian(t *testing.T) {
	t.Parallel()
	delta := 0.01
	for _, tc := range testcases {
		rm := NewRunningMedian()
		var mismatch bool
		results := make([]int, 0, len(tc.input))
		for i, v := range tc.input {
			result := rm.Median(v)
			if lower, upper := result-delta, result+delta; tc.expected[i] > upper || tc.expected[i] < lower {
				mismatch = true
				results = append(results, v)
			}
		}
		if mismatch {
			t.Errorf("Expected %v got %v", tc.expected, results)
		}
	}
}

func TestHeaps(t *testing.T) {
	t.Parallel()
	data := []int{1, 2, 3}
	min := &MinPQ{}
	max := &MaxPQ{}
	for _, v := range data {
		heap.Push(min, v)
		heap.Push(max, v)
	}
	if min.Len() != max.Len() {
		t.Errorf("Lengths should match %v and %v", min.Len(), max.Len())
	}
	for i := range data {
		x := heap.Pop(min)
		if x != data[i] {
			t.Errorf("Expected %v but got %v", data[i], x)
		}
	}
	for i := range data {
		x := heap.Pop(max)
		if x != data[len(data)-i-1] {
			t.Errorf("Expected %v but got %v", data[len(data)-i-1], x)
		}
	}
}

func BenchmarkRunningMedian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			rm := NewRunningMedian()
			for _, v := range tc.input {
				rm.Median(v)
			}
		}
	}
}

package day147

import (
	"reflect"
	"sort"
	"testing"
)

func TestReverse(t *testing.T) {
	t.Parallel()
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	Reverse(sort.IntSlice(data), 3, 8)
	expected := []int{1, 2, 3, 8, 7, 6, 5, 4, 9}
	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v got %v", expected, data)
	}
}

func TestSortRangeReverse(t *testing.T) {
	t.Parallel()
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	SortRange(sort.Reverse(sort.IntSlice(data)), 3, 8)
	expected := []int{1, 2, 3, 8, 7, 6, 5, 4, 9}
	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v got %v", expected, data)
	}
}

func TestSortRangeOutOfRange(t *testing.T) {
	t.Parallel()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Bad indicies should panic")
		}
	}()
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	SortRange(sort.Reverse(sort.IntSlice(data)), -1, 10)
}

func TestSortRangeBadRange(t *testing.T) {
	t.Parallel()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("i should be less than j")
		}
	}()
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	SortRange(sort.Reverse(sort.IntSlice(data)), 5, 3)
}

func BenchmarkReverse(b *testing.B) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		Reverse(sort.IntSlice(data), 3, 8)
	}
}

func BenchmarkSortRangeReverse(b *testing.B) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		SortRange(sort.Reverse(sort.IntSlice(data)), 3, 8)
	}
}

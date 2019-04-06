package day134

import (
	"testing"
)

func TestSparseArray(t *testing.T) {
	t.Parallel()
	arr := []int{0, 0, 0, 0, 9, 8, 7, 6, 0, 0}
	sa := NewSparseArray(arr)
	for i, val := range arr {
		if result, err := sa.Get(i); result != val || err != nil {
			t.Errorf("Expected (%v,%v) got (%v,%v)", val, nil, result, err)
		}
	}
	if err := sa.Set(0, 10); err != nil {
		t.Error("set shouldn't return an error", err)
	}
	if result, err := sa.Get(0); result != 10 || err != nil {
		t.Error("Set didn't work properly")
	}
	if err := sa.Set(0, 0); err != nil {
		t.Error("set shouldn't return an error", err)
	}
	for i, val := range arr {
		if result, err := sa.Get(i); result != val || err != nil {
			t.Errorf("Expected (%v,%v) got (%v,%v)", val, nil, result, err)
		}
	}
}

func TestSparseArrayNoInit(t *testing.T) {
	t.Parallel()
	var sa *SparseArray
	if err := sa.Set(0, 1); err == nil {
		t.Errorf("Set returns an error if called before Init")
	}
	if _, err := sa.Get(0); err == nil {
		t.Errorf("Get returns an error if called before Init")
	}
}

func TestSparseArrayOutOfRange(t *testing.T) {
	t.Parallel()
	sa := NewSparseArray([]int{1})
	if err := sa.Set(1, 1); err != ErrIndexOutOfRange() {
		t.Errorf("Index 1 is out of range")
	}
	if _, err := sa.Get(-1); err != ErrIndexOutOfRange() {
		t.Errorf("Index -1 is out of range")
	}
}

func BenchmarkSparseArray(b *testing.B) {
	arr := []int{0, 0, 0, 0, 9, 8, 7, 6, 0, 0}
	sa := NewSparseArray(arr)
	for i := 0; i < b.N; i++ {
		//nolint: errcheck
		sa.Set(0, 10)
		//nolint: errcheck
		sa.Get(0)
		//nolint: errcheck
		sa.Set(0, 0)

	}
}

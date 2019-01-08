package day137

import (
	"testing"
)

var testcases = []struct {
	ones map[int]bool
	size int
}{
	{map[int]bool{0: true, 5: true, 9: true}, 10},
}

func TestBitArray(t *testing.T) {
	t.Parallel()
	var ba BitArray
	for _, tc := range testcases {
		ba.Init(tc.size)
		for index, val := range tc.ones {
			if err := ba.Set(index, val); err != nil {
				t.Errorf("Sets errored but shouldn't")
			}
		}
		for i := 0; i < tc.size; i++ {
			if val, err := ba.Get(i); err != nil {
				t.Errorf("Error retrieving values that shouldn't %v", err)
			} else if tc.ones[i] != val {
				t.Errorf("Expected %v got %v for index %d", tc.ones[i], val, i)
			}
		}
	}
}

func TestBitArrayOutOfRange(t *testing.T) {
	t.Parallel()
	var ba BitArray
	ba.Init(10)
	if err := ba.Set(-1, true); err == nil {
		t.Errorf("-1 should be an invalid index")
	}
	if err := ba.Set(11, true); err == nil {
		t.Errorf("11 should be an invalid index when size is 10")
	}
	if _, err := ba.Get(-1); err == nil {
		t.Errorf("-1 should be an invalid index")
	}
	if _, err := ba.Get(11); err == nil {
		t.Errorf("11 should be an invalid index when size is 10")
	}
}

func TestBitArraySetZeros(t *testing.T) {
	t.Parallel()
	var ba BitArray
	ba.Init(10)
	for i := 0; i < 10; i++ {
		if bit, _ := ba.Get(i); bit {
			t.Errorf("Should not be set")
		}
	}
	ba.Set(0, false)
	ba.Set(5, false)
	ba.Set(9, false)
	for i := 0; i < 10; i++ {
		if bit, _ := ba.Get(i); bit {
			t.Errorf("Should not be set")
		}
	}
	ba.Set(0, false)
	ba.Set(5, false)
	ba.Set(9, false)
	for i := 0; i < 10; i++ {
		if bit, _ := ba.Get(i); bit {
			t.Errorf("Should not be set")
		}
	}
}

func BenchmarkBitArray(b *testing.B) {
	var ba BitArray
	ba.Init(10)
	for i := 0; i < b.N; i++ {
		ba.Set(i%10, i%2 == 0)
		ba.Get(i % 10)
	}
}

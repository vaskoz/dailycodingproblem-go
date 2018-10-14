package day52

import "testing"

func TestLRUCache(t *testing.T) {
	t.Parallel()
	lru := NewLRU(3)
	for i := 0; i < 6; i++ {
		lru.Set(i, i+1)
	}
	for i := 0; i < 3; i++ {
		if v := lru.Get(i); v != nil {
			t.Errorf("Expected these to be evicted")
		}
	}
	for i := 3; i < 6; i++ {
		if v := lru.Get(i); v != i+1 {
			t.Errorf("Expected the values to remain")
		}
	}
	for i := 3; i < 6; i++ {
		lru.Set(i, 10*i)
	}
	for i := 3; i < 6; i++ {
		if v := lru.Get(i); v != i*10 {
			t.Errorf("Expected the values to modify")
		}
	}
}

func BenchmarkLRUCache(b *testing.B) {
	lru := NewLRU(3)
	for bm := 0; bm < b.N; bm++ {
		for i := 0; i < 6; i++ {
			lru.Set(i, i+1)
		}
		for i := 0; i < 3; i++ {
			lru.Get(i)
		}
		for i := 3; i < 6; i++ {
			lru.Get(i)
		}
		for i := 3; i < 6; i++ {
			lru.Set(i, 10*i)
		}
		for i := 3; i < 6; i++ {
			lru.Get(i)
		}
	}
}

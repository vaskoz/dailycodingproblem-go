package day67

import "testing"

func TestLFUCacheOneEviction(t *testing.T) {
	t.Parallel()
	lfu := NewLFUCache(3)
	lfu.Set("foo", "bar")
	lfu.Set(1, 2)
	lfu.Set("bar", "baz")
	lfu.Get(1)
	lfu.Get("bar")
	lfu.Set(100, 1)
	if result := lfu.Get("foo"); result != nil {
		t.Errorf("'foo' should have been evicted at this point, got %v", result)
	}
	if result := lfu.Get(1); result == nil {
		t.Errorf("1 should remain")
	}
	if result := lfu.Get("bar"); result == nil {
		t.Errorf("'bar' should remain")
	}
	if result := lfu.Get(100); result == nil {
		t.Errorf("100 should remain")
	}
	lfu.Set(100, "baz")
	if result := lfu.Get(100); result.(string) != "baz" {
		t.Errorf("I expected baz")
	}
}

func TestLFUCacheDifferentAccessAmounts(t *testing.T) {
	t.Parallel()
	lfu := NewLFUCache(3)
	lfu.Set("foo", "bar")
	lfu.Set("foo", "bas")
	lfu.Set("foo", "bat")
	lfu.Set("foo", "bau")
	lfu.Set(1, 2)
	lfu.Set(1, 3)
	lfu.Set("bar", "baz")
	lfu.Set("willevictbar", "byebye")
	lfu.Set("willevictbar2", "byebye")
	if result := lfu.Get("foo"); result == nil {
		t.Errorf("Should remain foo: %v", result)
	}
	if result := lfu.Get(1); result == nil {
		t.Errorf("Shouldn't be nil, 1: %v", result)
	}
	if result := lfu.Get("bar"); result != nil {
		t.Errorf("Should be nil")
	}
	if result := lfu.Get("willevictbar"); result != nil {
		t.Errorf("Should be nil")
	}
	if result := lfu.Get("willevictbar2"); result == nil {
		t.Errorf("Should remain")
	}
}

func TestLFUCacheAccessGaps(t *testing.T) {
	t.Parallel()
	lfu := NewLFUCache(3)
	lfu.Set("foo", "bar")
	lfu.Set("foo", "bas")
	lfu.Set("foo", "bat")
	lfu.Set("foo", "bau")
	lfu.Set("another", "one")
	lfu.Set("last", "one")
	if result := lfu.Get("foo"); result != "bau" {
		t.Errorf("Expected 'bau' got %v", result)
	}
	if result := lfu.Get("another"); result != "one" {
		t.Errorf("Expected 'one' got %v", result)
	}
	if result := lfu.Get("last"); result != "one" {
		t.Errorf("Expected 'one' got %v", result)
	}
}

func BenchmarkLFUCache(b *testing.B) {
	lfu := NewLFUCache(3)
	for i := 0; i < b.N; i++ {
		lfu.Set("foo", "bar")
		lfu.Set(1, 2)
		lfu.Set("bar", "baz")
		lfu.Get(1)
		lfu.Get("bar")
		lfu.Set(100, 1)
	}
}

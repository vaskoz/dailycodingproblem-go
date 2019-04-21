package day242

import (
	"sync"
	"testing"
)

// Read-optimized
func TestReadOptimized(t *testing.T) {
	t.Parallel()
	input := [24]int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1100, 1200, 1300, 1400,
		1500, 1600, 1700, 1800, 1900, 2000, 2100, 2200, 2300, 2400}
	ds := NewReadOptimized(input)
	if sum := ds.Query(0, 23); sum != 30000 {
		t.Errorf("Expected 30000, got %v", sum)
	}
	ds.Update(20, 100)
	if sum := ds.Query(0, 23); sum != 30100 {
		t.Errorf("Expected 30100, got %v", sum)
	}
	ds.Update(20, 1000)
	if sum := ds.Query(20, 20); sum != 3200 {
		t.Errorf("Expected 3200, got %v", sum)
	}
}

func TestReadOptimizedConcurrently(t *testing.T) {
	t.Parallel()
	input := [24]int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1100, 1200, 1300, 1400,
		1500, 1600, 1700, 1800, 1900, 2000, 2100, 2200, 2300, 2400}
	ds := NewReadOptimized(input)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 24; i++ {
			ds.Update(i, 100)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 24; i++ {
			ds.Update(i, -200)
		}
		wg.Done()
	}()
	wg.Wait()
	if sum := ds.Query(0, 23); sum != 27600 {
		t.Errorf("Expected 27600, got %v", sum)
	}
}

func TestReadOptimizedInvalidInputUpdate(t *testing.T) {
	t.Parallel()
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("expected an error for bad hour value")
		}
	}()
	ds := NewReadOptimized([24]int{})
	ds.Update(24, 100)
}

func TestReadOptimizedInvalidInputQuery(t *testing.T) {
	t.Parallel()
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("expected an error for bad hour value")
		}
	}()
	ds := NewReadOptimized([24]int{})
	ds.Query(-1, 100)
}

func TestReadOptimizedQueryStartAfterEnd(t *testing.T) {
	t.Parallel()
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("expected an error for bad hour value")
		}
	}()
	ds := NewReadOptimized([24]int{})
	ds.Query(21, 20)
}

func BenchmarkReadOptimizedReads(b *testing.B) {
	input := [24]int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1100, 1200, 1300, 1400,
		1500, 1600, 1700, 1800, 1900, 2000, 2100, 2200, 2300, 2400}
	ds := NewReadOptimized(input)
	for i := 0; i < b.N; i++ {
		ds.Query(0, 23)
	}
}

func BenchmarkReadOptimizedWrites(b *testing.B) {
	input := [24]int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1100, 1200, 1300, 1400,
		1500, 1600, 1700, 1800, 1900, 2000, 2100, 2200, 2300, 2400}
	ds := NewReadOptimized(input)
	for i := 0; i < b.N; i++ {
		ds.Update(i%24, 1)
	}
}

// Write-optimized
func TestWriteOptimized(t *testing.T) {
	t.Parallel()
	input := [24]int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1100, 1200, 1300, 1400,
		1500, 1600, 1700, 1800, 1900, 2000, 2100, 2200, 2300, 2400}
	ds := NewWriteOptimized(input)
	if sum := ds.Query(0, 23); sum != 30000 {
		t.Errorf("Expected 30000, got %v", sum)
	}
	ds.Update(20, 100)
	if sum := ds.Query(0, 23); sum != 30100 {
		t.Errorf("Expected 30100, got %v", sum)
	}
	ds.Update(20, 1000)
	if sum := ds.Query(20, 20); sum != 3200 {
		t.Errorf("Expected 3200, got %v", sum)
	}
}

func TestWriteOptimizedConcurrently(t *testing.T) {
	t.Parallel()
	input := [24]int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1100, 1200, 1300, 1400,
		1500, 1600, 1700, 1800, 1900, 2000, 2100, 2200, 2300, 2400}
	ds := NewWriteOptimized(input)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 24; i++ {
			ds.Update(i, 100)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 24; i++ {
			ds.Update(i, -200)
		}
		wg.Done()
	}()
	wg.Wait()
	if sum := ds.Query(0, 23); sum != 27600 {
		t.Errorf("Expected 27600, got %v", sum)
	}
}

func TestWriteOptimizedInvalidInputUpdate(t *testing.T) {
	t.Parallel()
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("expected an error for bad hour value")
		}
	}()
	ds := NewWriteOptimized([24]int{})
	ds.Update(24, 100)
}

func TestWriteOptimizedInvalidInputQuery(t *testing.T) {
	t.Parallel()
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("expected an error for bad hour value")
		}
	}()
	ds := NewWriteOptimized([24]int{})
	ds.Query(-1, 100)
}

func TestWriteOptimizedQueryStartAfterEnd(t *testing.T) {
	t.Parallel()
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("expected an error for bad hour value")
		}
	}()
	ds := NewWriteOptimized([24]int{})
	ds.Query(21, 20)
}

func BenchmarkWriteOptimizedReads(b *testing.B) {
	input := [24]int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1100, 1200, 1300, 1400,
		1500, 1600, 1700, 1800, 1900, 2000, 2100, 2200, 2300, 2400}
	ds := NewWriteOptimized(input)
	for i := 0; i < b.N; i++ {
		ds.Query(0, 23)
	}
}

func BenchmarkWriteOptimizedWrites(b *testing.B) {
	input := [24]int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1100, 1200, 1300, 1400,
		1500, 1600, 1700, 1800, 1900, 2000, 2100, 2200, 2300, 2400}
	ds := NewWriteOptimized(input)
	for i := 0; i < b.N; i++ {
		ds.Update(i%24, 1)
	}
}

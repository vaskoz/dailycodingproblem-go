package day132

import "testing"

func TestHitCounter(t *testing.T) {
	t.Parallel()
	hc := NewHitCounter()
	hc.Record(10)
	hc.Record(15)
	hc.Record(20)
	hc.Record(30)
	if total := hc.Total(); total != 4 {
		t.Errorf("Total should be 4 but got %v", total)
	}
	if r := hc.Range(0, 10); r != 1 {
		t.Errorf("Expected 1 got %v", r)
	}
	if r := hc.Range(10, 30); r != 4 {
		t.Errorf("Expected 4 got %v", r)
	}
	if r := hc.Range(11, 29); r != 2 {
		t.Errorf("Expected 2 got %v", r)
	}
}

func BenchmarkHitCounter(b *testing.B) {
	hc := NewHitCounter()
	hc.Record(10)
	hc.Record(15)
	hc.Record(20)
	hc.Record(30)

	for i := 0; i < b.N; i++ {

		hc.Total()
		hc.Range(0, 10)
		hc.Range(10, 30)
		hc.Range(11, 29)
	}
}

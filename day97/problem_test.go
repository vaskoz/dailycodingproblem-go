package day97

import "testing"

func TestTimeMap(t *testing.T) {
	t.Parallel()
	d := NewTimeMap()
	d.Set(1, 1, 0) // set key 1 to value 1 at time 0
	d.Set(1, 2, 2) // set key 1 to value 2 at time 2
	if result := d.Get(1, 1); result != 1 {
		t.Errorf("get key 1 at time 1 should be 1")
	}
	if result := d.Get(1, 3); result != 2 {
		t.Errorf("get key 1 at time 3 should be 2")
	}
	d.Set(1, 1, 5) // set key 1 to value 1 at time 5
	if result := d.Get(1, 0); result != 1 {
		t.Errorf("get key 1 at time 0 should be 1")
	}
	if result := d.Get(1, 10); result != 1 {
		t.Errorf("get key 1 at time 10 should be 1")
	}
	d.Set(1, 1, 0) // set key 1 to value 1 at time 0
	d.Set(1, 2, 0) // set key 1 to value 2 at time 0
	if result := d.Get(1, 0); result != 2 {
		t.Errorf("get key 1 at time 0 should be 2")
	}
	if result := d.Get(2, 0); result != nil {
		t.Errorf("unused keys should return nil")
	}
	if result := d.Get(1, -1); result != nil {
		t.Errorf("No value should be set back in time.")
	}
}

func BenchmarkTimeMap(b *testing.B) {
	d := NewTimeMap()
	d.Set(1, 1, 0) // set key 1 to value 1 at time 0
	d.Set(1, 2, 2) // set key 1 to value 2 at time 2
	d.Set(1, 1, 5) // set key 1 to value 1 at time 5
	for i := 0; i < b.N; i++ {
		d.Get(1, 1)
		d.Get(1, 3)
		d.Get(1, 0)
		d.Get(1, 10)
		d.Get(1, 0)
		d.Get(2, 0)
		d.Get(1, -1)
	}
}

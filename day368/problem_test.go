package day368

import "testing"

func TestKeyValueStore(t *testing.T) {
	t.Parallel()

	kv := NewKeyValueStore()

	if _, err := kv.Get(1); err == nil {
		t.Errorf("Expected an error getting 1 on empty KV store")
	}

	kv.Update(1, 1)
	kv.Update(2, 1)

	if v, err := kv.Get(1); v != 1 || err != nil {
		t.Errorf("Expected (%v,%v), got (%v,%v)", 1, nil, v, err)
	}

	if v, err := kv.Get(2); v != 1 || err != nil {
		t.Errorf("Expected (%v,%v), got (%v,%v)", 2, nil, v, err)
	}

	kv.Update(1, 2)
	kv.Update(2, 2)

	if v, err := kv.Get(1); v != 2 || err != nil {
		t.Errorf("Expected (%v,%v), got (%v,%v)", 2, nil, v, err)
	}

	if v, err := kv.MaxKey(2); v != 2 || err != nil {
		t.Errorf("Expected (%v,%v), got (%v,%v)", 2, nil, v, err)
	}

	if _, err := kv.MaxKey(1); err == nil {
		t.Errorf("Expected an error getting MaxKey 1")
	}
}

func BenchmarkKeyValueStore(b *testing.B) {
	kv := NewKeyValueStore()

	for i := 0; i < b.N; i++ {
		kv.Update(1, 1)
		kv.Update(2, 1)
		kv.Get(1)    // nolint
		kv.MaxKey(1) // nolint
		kv.MaxKey(2) // nolint
		kv.Update(1, 2)
		kv.Update(2, 2)
		kv.MaxKey(2) // nolint
	}
}

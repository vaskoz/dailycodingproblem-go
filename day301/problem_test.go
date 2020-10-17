package day301

import (
	"testing"

	"github.com/algds/bloom"
)

func TestBloomFilter(t *testing.T) {
	t.Parallel()

	f1 := func(d interface{}) uint {
		s := d.(string)

		return uint(s[0])
	}
	f2 := func(d interface{}) uint {
		s := d.(string)

		return uint(s[1])
	}

	blf := NewBloomFilter(100, []bloom.Hash{f1, f2})

	blf.Add("foo")
	blf.Add("bar")

	if !blf.Check("foo") {
		t.Errorf("foo should be found")
	}

	if !blf.Check("bar") {
		t.Errorf("bar should be found")
	}

	if blf.Check("xyz") {
		t.Errorf("xyz should not be found")
	}
}

func BenchmarkBloomFilter(b *testing.B) {
	f1 := func(d interface{}) uint {
		s := d.(string)

		return uint(s[0])
	}
	f2 := func(d interface{}) uint {
		s := d.(string)

		return uint(s[1])
	}

	blf := NewBloomFilter(100, []bloom.Hash{f1, f2})

	for i := 0; i < b.N; i++ {
		blf.Add("foo")
		blf.Add("bar")
		blf.Check("foo")
		blf.Check("bar")
		blf.Check("xyz")
	}
}

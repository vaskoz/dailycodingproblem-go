package day105

import (
	"testing"
	"time"
)

func TestDebounced(t *testing.T) {
	t.Parallel()
	x := 0
	f := func() {
		x++
	}
	dbf := Debounced(f, 500)
	for i := 0; i < 10; i++ {
		dbf()
	}
	if x != 0 {
		t.Errorf("Expected x not to change since it's debounced for 500ms")
	}
	time.Sleep(500 * time.Millisecond)
	for i := 0; i < 10; i++ {
		dbf()
	}
	if x != 10 {
		t.Errorf("Expected x to be incremented 10 times")
	}
}

package day59

import (
	"reflect"
	"testing"
)

func TestRsyncUpdates(t *testing.T) {
	t.Parallel()
	f1 := File("abcdefg")
	f2 := File("xyz123")
	src := &Rsync{}
	target := &Rsync{}
	src.AddTarget(target)
	src.AddFile(f1)
	src.AddFile(f2)

	f1a := File("abczefg")
	src.UpdateFile(0, f1a)
	if f := target.GetFile(0); !reflect.DeepEqual(f, f1a) {
		t.Errorf("f1 doesn't match inside the target. %v, %v", f, f1a)
	}
	f2a := File("bbbbb")
	src.UpdateFile(1, f2a)
	if f := target.GetFile(1); !reflect.DeepEqual(f, f2a) {
		t.Errorf("f2 doesn't match inside the target. %v, %v", f, f2a)
	}
	f2b := File("cccccccccccc")
	src.UpdateFile(1, f2b)
	if f := target.GetFile(1); !reflect.DeepEqual(f, f2b) {
		t.Errorf("f2 doesn't match inside the target. %v, %v", f, f2b)
	}
}

func TestRsync(t *testing.T) {
	t.Parallel()
	f1 := File("abcdefg")
	f2 := File("xyz123")
	src := &Rsync{}
	target := &Rsync{}
	src.AddTarget(target)
	src.AddFile(f1)
	src.AddFile(f2)

	if f := target.GetFile(0); !reflect.DeepEqual(f, f1) {
		t.Errorf("f1 doesn't match inside the target. %v, %v", f, f1)
	}
	if f := target.GetFile(1); !reflect.DeepEqual(f, f2) {
		t.Errorf("f2 doesn't match inside the target. %v, %v", f, f2)
	}
}

func TestRsyncBadInput(t *testing.T) {
	t.Parallel()
	f1 := File("abcdefg")
	f2 := File("xyz123")
	src := &Rsync{}
	target := &Rsync{}
	src.AddTarget(target)
	src.AddFile(f1)
	src.AddFile(f2)

	src.DeleteFile(0)
	src.DeleteFile(1)
	if f := target.GetFile(0); f != nil {
		t.Errorf("GetFile(0) should be nil. Got %v", f)
	}
	if f := target.GetFile(1); f != nil {
		t.Errorf("GetFile(1) should be nil. Got %v", f)
	}
	if err := target.DeleteFile(2); err == nil {
		t.Error("DeleteFile(2) should be an error")
	}
	if f := target.GetFile(5); f != nil {
		t.Errorf("GetFile(5) should be nil. Got %v", f)
	}
	if err := target.UpdateFile(5, f1); err == nil {
		t.Error("UpdateFile(5) should error")
	}
}

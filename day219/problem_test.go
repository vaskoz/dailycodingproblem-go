package day219

import "testing"

func TestConnectFourVerticalAlternate(t *testing.T) {
	t.Parallel()
	var cfb ConnectFourBoard
	for i := 0; i < 3; i++ {
		if err := cfb.Move(Red, 0); err != nil {
			t.Errorf("expected nil got %v", err)
		}
		if p := cfb.Winner(); p != None {
			t.Errorf("expected no winner got %v", p)
		}
		if err := cfb.Move(Black, 0); err != nil {
			t.Errorf("expected nil, got %v", err)
		}
		if p := cfb.Winner(); p != None {
			t.Errorf("expected no winner got %v", p)
		}
	}
	if err := cfb.Move(Red, 0); err == nil {
		t.Errorf("expected an error inserting into a full column")
	}
	if err := cfb.Move(None, 1); err == nil {
		t.Errorf("expected an error trying to insert non-red non-black")
	}
	if err := cfb.Move(Black+Red, 1); err == nil {
		t.Errorf("expected an error trying to insert non-red non-black")
	}
}

func TestConnectFourVerticalRedWins(t *testing.T) {
	t.Parallel()
	var cfb ConnectFourBoard
	for i := 0; i < 3; i++ {
		if err := cfb.Move(Red, 0); err != nil {
			t.Errorf("expected nil got %v", err)
		}
		if p := cfb.Winner(); p != None {
			t.Errorf("expected no winner got %v", p)
		}
		if err := cfb.Move(Black, 1); err != nil {
			t.Errorf("expected nil, got %v", err)
		}
		if p := cfb.Winner(); p != None {
			t.Errorf("expected no winner got %v", p)
		}
	}
	if err := cfb.Move(Red, 0); err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	if p := cfb.Winner(); p != Red {
		t.Errorf("expected Red winner got %v", p)
	}
}

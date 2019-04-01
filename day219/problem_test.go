package day219

import (
	"testing"
)

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

func TestConnectFourHorizontalBlackWins(t *testing.T) {
	t.Parallel()
	var cfb ConnectFourBoard
	for i := Column(0); i < 3; i++ {
		if err := cfb.Move(Black, i); err != nil {
			t.Errorf("expected nil got %v", err)
		}
		if p := cfb.Winner(); p != None {
			t.Errorf("expected no winner got %v", p)
		}
		if err := cfb.Move(Red, i); err != nil {
			t.Errorf("expected nil, got %v", err)
		}
		if p := cfb.Winner(); p != None {
			t.Errorf("expected no winner got %v", p)
		}
	}
	if err := cfb.Move(Black, 3); err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	if p := cfb.Winner(); p != Black {
		t.Errorf("expected Black winner got %v", p)
	}
}

func TestConnectFourDiagonalDownRightRedWins(t *testing.T) {
	t.Parallel()
	var cfb ConnectFourBoard
	for i := 0; i < 3; i++ {
		if err := cfb.Move(Black, Column(0)); err != nil {
			t.Errorf("expected nil got %v", err)
		}
	}
	if err := cfb.Move(Red, Column(0)); err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	for i := 0; i < 2; i++ {
		if err := cfb.Move(Black, Column(1)); err != nil {
			t.Errorf("expected nil got %v", err)
		}
	}
	if err := cfb.Move(Red, Column(1)); err != nil {
		t.Errorf("expected nil got %v", err)
	}
	if err := cfb.Move(Black, Column(2)); err != nil {
		t.Errorf("expected nil got %v", err)
	}
	if err := cfb.Move(Red, Column(2)); err != nil {
		t.Errorf("expected nil got %v", err)
	}
	if p := cfb.Winner(); p != None {
		t.Errorf("expected no winner got %v", p)
	}
	if err := cfb.Move(Red, Column(3)); err != nil {
		t.Errorf("expected nil got %v", err)
	}
	if p := cfb.Winner(); p != Red {
		t.Errorf("expected Red winner got %v", p)
	}
}

func TestConnectFourDiagonalUpRightRedWins(t *testing.T) {
	t.Parallel()
	var cfb ConnectFourBoard
	if err := cfb.Move(Red, Column(0)); err != nil {
		t.Errorf("expected nil got %v", err)
	}
	if err := cfb.Move(Black, Column(1)); err != nil {
		t.Errorf("expected nil got %v", err)
	}
	if err := cfb.Move(Red, Column(1)); err != nil {
		t.Errorf("expected nil got %v", err)
	}
	for i := 0; i < 2; i++ {
		if err := cfb.Move(Black, Column(2)); err != nil {
			t.Errorf("expected nil got %v", err)
		}
		if p := cfb.Winner(); p != None {
			t.Errorf("expected no winner got %v", p)
		}
	}
	if err := cfb.Move(Red, Column(2)); err != nil {
		t.Errorf("expected nil got %v", err)
	}
	if p := cfb.Winner(); p != None {
		t.Errorf("expected no winner got %v", p)
	}
	for i := 0; i < 3; i++ {
		if err := cfb.Move(Black, Column(3)); err != nil {
			t.Errorf("expected nil got %v", err)
		}
		if p := cfb.Winner(); p != None {
			t.Errorf("expected no winner got %v", p)
		}
	}
	if err := cfb.Move(Red, Column(3)); err != nil {
		t.Errorf("expected nil got %v", err)
	}
	if p := cfb.Winner(); p != Red {
		t.Errorf("expected Red winner got %v", p)
	}
}

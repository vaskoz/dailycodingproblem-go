package day239

import "testing"

func TestValidUnlockKeypadNumber(t *testing.T) {
	t.Parallel()
	if result := ValidUnlockKeypadNumber(); result != 152316 {
		t.Errorf("Expected 152316, got %v", result)
	}
}

func BenchmarkValidUnlockKeypadNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ValidUnlockKeypadNumber()
	}
}

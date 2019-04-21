package day229

import "testing"

func TestSnakesAndLaddersFewestTurns(t *testing.T) {
	t.Parallel()
	fewestTurns := SnakesAndLaddersFewestTurns(
		[]Jump{{1, 38}, {4, 14}, {9, 31}, {21, 42}, {28, 84}, {36, 44}, {51, 67}, {71, 91}, {80, 100}},
		[]Jump{{16, 6}, {48, 26}, {49, 11}, {56, 53}, {62, 19}, {64, 60}, {87, 24}, {93, 73}, {95, 75}, {98, 78}},
		6,
	)
	if fewestTurns != 7 {
		t.Errorf("on the original game board best is 7, got %v", fewestTurns)
	}
}

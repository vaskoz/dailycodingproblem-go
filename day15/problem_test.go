package day15

import (
	"math/rand"
	"testing"
)

func TestUniformlyRandom(t *testing.T) {
	t.Parallel()
	randomizer = rand.New(rand.NewSource(99)) // use a fixed seed for testing.
	expected := 58                            // because of the constant random seed
	input := make(chan interface{})
	go func() {
		for i := 0; i < 100; i++ {
			input <- i
		}
		close(input)
	}()
	if result := UniformlyRandom(input); result != 58 {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func BenchmarkUniformlyRandom(b *testing.B) {
	randomizer = rand.New(rand.NewSource(99)) // use a fixed seed for testing.
	for i := 0; i < b.N; i++ {
		input := make(chan interface{})
		go func() {
			for i := 0; i < 100; i++ {
				input <- i
			}
			close(input)
		}()
		UniformlyRandom(input)
	}
}

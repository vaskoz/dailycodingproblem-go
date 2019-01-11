package day141

import (
	"testing"
)

func TestThreeStacks(t *testing.T) {
	t.Parallel()
	var stack ThreeStacks
	for i := 0; i < 100; i++ {
		stack.Push(i, i%3)
	}
	for i := 99; i >= 0; i-- {
		if v, err := stack.Pop(i % 3); err != nil || v != i {
			t.Errorf("Expected (%v,nil) got (%v,%v)", i, v, err)
		}
	}
}

func TestThreeStacksErrors(t *testing.T) {
	t.Parallel()
	var stack ThreeStacks

	if err := stack.Push("foobar", 3); err != ErrInvalidStackID() {
		t.Errorf("Expected an invalid stack id error for Push")
	}
	if _, err := stack.Pop(-1); err != ErrInvalidStackID() {
		t.Errorf("Expected an invalid stack id error for Pop")
	}
	if _, err := stack.Pop(0); err != ErrEmptyStack() {
		t.Errorf("Expected an empty stack error for Pop")
	}
}

func TestThreeStacksDefrag(t *testing.T) {
	t.Parallel()
	var stack ThreeStacks
	for i := 0; i < 100; i++ {
		stack.Push(i, i%3)
	}
	for i := 99; i >= 0; i-- {
		stack.Pop(i % 3)
	}
	if len(stack.data) > 20 {
		t.Errorf("The list should have reduced due to defrag")
	}
}

func BenchmarkThreeStacks(b *testing.B) {
	var stack ThreeStacks
	for i := 0; i < 10000; i++ {
		stack.Push(i, i%3)
	}
	for i := 0; i < b.N; i++ {
		stack.Push(i, i%3)
		for j := 0; j < 30; j++ {
			stack.Pop(j % 3)
		}
	}
}

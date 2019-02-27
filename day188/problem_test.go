package day188

import "testing"

func TestMakeFunctionsOriginal(t *testing.T) {
	t.Parallel()
	funcs := MakeFunctionsOriginal()
	for _, f := range funcs {
		if result := f(); result != 3 {
			t.Errorf("Expected every response to be 3. Got %v", result)
		}
	}
}

func TestMakeFunctionsCorrect(t *testing.T) {
	t.Parallel()
	funcs := MakeFunctionsCorrect()
	for i, f := range funcs {
		if result := f(); result != i+1 {
			t.Errorf("Expected %v, got %v", i+1, result)
		}
	}
}

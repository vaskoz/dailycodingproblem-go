package day174

import "testing"

func TestAskPolymorph(t *testing.T) {
	t.Parallel()
	var sloth Sloth
	var industrious Industrious
	if msg := AskPolymorph(sloth); msg != slothMsg {
		t.Errorf("Got %v, expected %v", msg, slothMsg)
	}
	if msg := AskPolymorph(industrious); msg != industriousMsg {
		t.Errorf("Got %v, expected %v", msg, industriousMsg)
	}
}

func BenchmarkAskPolymorph(b *testing.B) {
	var sloth Sloth
	var industrious Industrious
	for i := 0; i < b.N; i++ {
		AskPolymorph(sloth)
		AskPolymorph(industrious)
	}
}

package day55

import "testing"

var urls = []struct {
	url, short string
}{
	{"http://google.com", "AAAAAA"},
	{"http://yahoo.com", "AAAAAB"},
	{"http://github.com", "AAAAAC"},
}

func TestInMemoryURLShortener(t *testing.T) {
	t.Parallel()
	us := NewInMemoryURLShortener()
	for _, tc := range urls {
		if short := us.Shorten(tc.url); short != tc.short {
			t.Errorf("Expected %s got %s", tc.short, short)
		}
	}
	for _, tc := range urls {
		if url := us.Restore(tc.short); url != tc.url {
			t.Errorf("Expected %s got %s", tc.url, url)
		}
	}
}

func BenchmarkInMemoryURLShortener(b *testing.B) {
	us := NewInMemoryURLShortener()
	for i := 0; i < b.N; i++ {
		for _, tc := range urls {
			us.Shorten(tc.url)
		}
	}
}

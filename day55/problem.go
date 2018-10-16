package day55

const base64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

var (
	b64pos = make(map[rune]int)
	posb64 = make(map[int]rune)
)

func init() {
	for i, r := range base64 {
		b64pos[r] = i
		posb64[i] = r
	}
}

// URLShortener represents a URL shortener application.
type URLShortener interface {
	Shorten(url string) string
	Restore(short string) string
}

type inMemoryURLShortener struct {
	shorten map[string]string
	restore map[string]string
	next    uint64
}

// NewInMemoryURLShortener returns an in-memory URLShortener.
func NewInMemoryURLShortener() URLShortener {
	return &inMemoryURLShortener{
		shorten: make(map[string]string),
		restore: make(map[string]string),
	}
}

func (imus *inMemoryURLShortener) Shorten(url string) string {
	k := key(imus.next)
	imus.shorten[url] = k
	imus.restore[k] = url
	imus.next++
	return k
}

func (imus *inMemoryURLShortener) Restore(short string) string {
	return imus.restore[short]
}

func key(index uint64) string {
	result := make([]rune, 6)
	i := 5
	for index != 0 {
		x := int(index % 64)
		result[i] = posb64[x]
		index /= 64
		i--
	}
	for i >= 0 {
		result[i] = posb64[0]
		i--
	}
	return string(result)
}

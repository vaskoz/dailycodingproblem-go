package day261

import (
	"container/heap"
	"errors"
	"strings"
)

var errTranslate = errors.New("could not be translated")

// Huffman is a huffman tree that can encode/decode.
type Huffman interface {
	Encode(string) (string, error)
	Decode(string) (string, error)
}

type huffman struct {
	Freq        int
	Leaf        bool
	R           rune
	Left, Right *huffman
}

func dfsBuildCodeMap(h *huffman, codeMap map[rune]string, path string) map[rune]string {
	if h == nil {
		return codeMap
	} else if h.Leaf {
		codeMap[h.R] = path
		return codeMap
	}
	codeMap = dfsBuildCodeMap(h.Left, codeMap, path+"0")
	codeMap = dfsBuildCodeMap(h.Right, codeMap, path+"1")
	return codeMap
}

func buildCodeMap(h *huffman) map[rune]string {
	return dfsBuildCodeMap(h, make(map[rune]string), "")
}

func (h *huffman) Encode(s string) (string, error) {
	var sb strings.Builder
	codeMap := buildCodeMap(h)
	for _, r := range s {
		prefix, found := codeMap[r]
		if !found {
			return "", errTranslate
		}
		sb.WriteString(prefix)
	}
	return sb.String(), nil
}

func (h *huffman) Decode(s string) (string, error) {
	var sb strings.Builder
	root := h
	for _, r := range s {
		switch r {
		case '0':
			if root.Leaf {
				sb.WriteRune(root.R)
				root = h
			}
			root = root.Left

		case '1':
			if root.Leaf {
				sb.WriteRune(root.R)
				root = h
			}
			root = root.Right
		}
	}
	if root.Leaf {
		sb.WriteRune(root.R)
		return sb.String(), nil
	}
	return "", errTranslate
}

type minHeap []*huffman

func (mh minHeap) Len() int { return len(mh) }

func (mh minHeap) Less(i, j int) bool {
	return mh[i].Freq < mh[j].Freq
}

func (mh minHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *minHeap) Push(x interface{}) {
	item := x.(*huffman)
	*mh = append(*mh, item)
}

func (mh *minHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	item := old[n-1]
	*mh = old[0 : n-1]
	return item
}

// NewHuffman returns a new Huffman tree.
func NewHuffman(freq map[rune]int) Huffman {
	mh := make(minHeap, 0, len(freq))
	for r, count := range freq {
		heap.Push(&mh, &huffman{Freq: count, R: r, Leaf: true})
	}
	for mh.Len() > 1 {
		one := heap.Pop(&mh).(*huffman)
		two := heap.Pop(&mh).(*huffman)
		parent := &huffman{Freq: one.Freq + two.Freq, Left: one, Right: two}
		mh.Push(parent)
	}
	return mh.Pop().(*huffman)
}

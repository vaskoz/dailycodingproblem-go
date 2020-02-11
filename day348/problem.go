package day348

// TernarySearchTree implements this tree:
// https://en.wikipedia.org/wiki/Ternary_search_tree
type TernarySearchTree struct {
	letter              rune
	left, middle, right *TernarySearchTree
}

// Insert adds a word to the TernarySearchTree.
func (tst *TernarySearchTree) Insert(word string) {
	if word == "" {
		return
	}

	switch remaining, r := len(word), rune(word[0]); {
	case tst.letter == 0 || tst.letter == r:
		tst.letter = r
		if remaining != 1 && tst.middle == nil {
			tst.middle = &TernarySearchTree{}
		}

		tst = tst.middle
		tst.Insert(word[1:])
	case tst.letter > r:
		if tst.left == nil {
			tst.left = &TernarySearchTree{}
		}

		tst = tst.left
		tst.Insert(word)
	case tst.letter < r:
		if tst.right == nil {
			tst.right = &TernarySearchTree{}
		}

		tst = tst.right
		tst.Insert(word)
	}
}

// Search tries to find the word in the TernarySearchTree.
func (tst *TernarySearchTree) Search(word string) bool {
	if word == "" {
		return true
	} else if tst == nil {
		return false
	}

	switch r := rune(word[0]); {
	case tst.letter == r:
		tst = tst.middle
		return tst.Search(word[1:])
	case tst.letter > r:
		tst = tst.left
		return tst.Search(word)
	default: // case tst.letter < r:
		tst = tst.right
		return tst.Search(word)
	}
}

package day196

// BinaryTree is a binary tree of integers.
type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

// MostFrequentSubtreeSum returns the most frequent subtree sum.
// Runs in linear time.
func MostFrequentSubtreeSum(head *BinaryTree) int {
	freq := make(map[int]int)
	subtreeFreqCount(head, freq)
	var result, mostFreq int
	for sum, count := range freq {
		if count > mostFreq {
			mostFreq = count
			result = sum
		}
	}
	return result
}

func subtreeFreqCount(head *BinaryTree, freq map[int]int) int {
	if head == nil {
		return 0
	}
	left := subtreeFreqCount(head.Left, freq)
	right := subtreeFreqCount(head.Right, freq)
	sum := left + right + head.Value
	freq[sum]++
	return sum
}

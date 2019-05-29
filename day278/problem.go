package day278

// BST is an integer binary search tree.
type BST struct {
	Value       int
	Left, Right *BST
}

func ConstructAllPossibleBSTs(n int) []*BST {
	return constructAllPossibleBSTs(1, n)
}

func constructAllPossibleBSTs(start, end int) []*BST {
	var result []*BST
	if start > end {
		return []*BST{nil}
	}
	for i := start; i < end+1; i++ {
		left := constructAllPossibleBSTs(start, i-1)
		right := constructAllPossibleBSTs(i+1, end)
		for j := range left {
			l := left[j]
			for k := range right {
				r := right[k]
				node := &BST{Value: i}
				node.Left = l
				node.Right = r
				result = append(result, node)
			}
		}
	}
	return result
}
